package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"github.com/tkanos/gonfig"
	"log"
	"net/http"
	"time"
	"ttnmapper-postgres-insert-raw/pkg/database"
	"ttnmapper-postgres-insert-raw/pkg/queues"
	"ttnmapper-postgres-insert-raw/pkg/types"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

type Configuration struct {
	AmqpHost                 string `env:"AMQP_HOST"`
	AmqpPort                 string `env:"AMQP_PORT"`
	AmqpUser                 string `env:"AMQP_USER"`
	AmqpPassword             string `env:"AMQP_PASSWORD"`
	AmqpExchangeRawData      string `env:"AMQP_EXCHANGE_RAW"`
	AmqpQueue                string `env:"AMQP_QUEUE"`
	AmqpExchangeInsertedData string `env:"AMQP_EXCHANGE_INSERTED"`

	PostgresHost          string `env:"POSTGRES_HOST"`
	PostgresPort          string `env:"POSTGRES_PORT"`
	PostgresUser          string `env:"POSTGRES_USER"`
	PostgresPassword      string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase      string `env:"POSTGRES_DATABASE"`
	PostgresDebugLog      bool   `env:"POSTGRES_DEBUG_LOG"`
	PostgresInsertThreads int    `env:"POSTGRES_INSERT_THREADS"`

	PrometheusPort string `env:"PROMETHEUS_PORT"`
}

var myConfiguration = Configuration{
	AmqpHost:                 "localhost",
	AmqpPort:                 "5672",
	AmqpUser:                 "user",
	AmqpPassword:             "password",
	AmqpExchangeRawData:      "new_packets",
	AmqpQueue:                "postgres_insert_raw",
	AmqpExchangeInsertedData: "inserted_data",

	PostgresHost:          "localhost",
	PostgresPort:          "5432",
	PostgresUser:          "username",
	PostgresPassword:      "password",
	PostgresDatabase:      "database",
	PostgresDebugLog:      false,
	PostgresInsertThreads: 1,

	PrometheusPort: "9100",
}

var (
	dbInserts = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ttnmapper_postgres_inserts_raw_count",
		Help: "The total number of packets inserted into the raw table",
	})
	invalidLocations = promauto.NewCounter(prometheus.CounterOpts{
		Name: "ttnmapper_postgres_invalid_locations",
		Help: "The total number of packets ignored due to location on null island",
	})

	insertDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "ttnmapper_postgres_inserts_raw_duration",
		Help:    "How long the processing and insert of a packet takes",
		Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.5, 2, 5, 10, 100, 1000, 10000},
	})

	messageChannel  = make(chan amqp.Delivery)
	insertedChannel = make(chan types.TtnMapperUplinkMessage)
)

func main() {

	err := gonfig.GetConf("conf.json", &myConfiguration)
	if err != nil {
		log.Println(err)
	}

	log.Printf("[Configuration]\n%s\n", utils.PrettyPrint(myConfiguration)) // output: [UserA, UserB]

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+myConfiguration.PrometheusPort, nil)
		if err != nil {
			log.Print(err.Error())
		}
	}()

	log.Println("Init database")
	databaseContext := database.DatabaseContext{
		Host:     myConfiguration.PostgresHost,
		Port:     myConfiguration.PostgresPort,
		User:     myConfiguration.PostgresUser,
		Database: myConfiguration.PostgresDatabase,
		Password: myConfiguration.PostgresPassword,
		DebugLog: myConfiguration.PostgresDebugLog,
	}
	databaseContext.Init()

	// Start threads to handle Postgres inserts
	log.Println("Starting database insert threads")
	for i := 0; i < myConfiguration.PostgresInsertThreads; i++ {
		go insertToPostgres(i + 1)
	}

	log.Println("Starting AMQP publish")
	queueCredentials := queues.QueueCredentials{
		User:     myConfiguration.AmqpUser,
		Password: myConfiguration.AmqpPassword,
		Host:     myConfiguration.AmqpHost,
		Port:     myConfiguration.AmqpPort,
	}
	publishQueueContext := queues.PublishContext{
		QueueCredentials: queueCredentials,
		Exchange:         myConfiguration.AmqpExchangeInsertedData,
		TxChannel:        insertedChannel,
	}
	go publishQueueContext.Publish()

	log.Println("Starting AMQP subscribe")
	subscribeQueueContext := queues.SubscribeContext{
		QueueCredentials: queueCredentials,
		Exchange:         myConfiguration.AmqpExchangeRawData,
		Queue:            myConfiguration.AmqpQueue,
		RxChannel:        messageChannel,
	}
	go subscribeQueueContext.Subscribe()

	forever := make(chan bool)
	<-forever
}

func insertToPostgres(thread int) {
	// Wait for a message and insert it into Postgres
	for d := range messageChannel {
		log.Printf("[%d][p] Processing packet", thread)

		// The message from amqp is a json string. Unmarshal to ttnmapper uplink struct
		var message types.TtnMapperUplinkMessage
		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Printf("[%d][p] "+err.Error(), thread)
			d.Ack(false)
			continue
		}

		// If coordinates are invalid, do not store - even if it's an experiment
		if message.Latitude == 0 && message.Longitude == 0 {
			log.Printf("[%d][p] Null island %s - %s - %s - %s", thread, message.NetworkId, message.AppID, message.DevID, message.UserId)
			invalidLocations.Inc()
			d.Ack(false)
			continue
		}

		// Iterate gateways. We store it flat in the database
		for _, gateway := range message.Gateways {
			gatewayStart := time.Now()

			// Copy required fields in correct format into a database row struct
			entry, err := database.UplinkMessageToPacket(message, gateway)
			if err != nil {
				log.Printf(err.Error())
				continue
			}

			// Insert into database
			err = database.InsertEntry(&entry)
			if err == nil {
				log.Printf("[%d][p] Inserted entry id=", thread, entry.ID)
				dbInserts.Inc()
			} else {
				log.Println(utils.PrettyPrint(entry))
				log.Print("[%d][p] PG Insert", thread)
				utils.FailOnError(err, "")
			}

			// Prometheus stats
			gatewayElapsed := time.Since(gatewayStart)
			insertDuration.Observe(float64(gatewayElapsed.Nanoseconds()) / 1000.0 / 1000.0) //nanoseconds to milliseconds
		}

		// If we get here all inserts were successful. Otherwise we would have quit.
		d.Ack(false)

		insertedChannel <- message
	}

	log.Fatal("Messages channel closed")
}
