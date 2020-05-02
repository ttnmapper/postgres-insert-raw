package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/streadway/amqp"
	"github.com/tkanos/gonfig"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
	"ttnmapper-postgres-insert-raw/types"
)

var messageChannel = make(chan amqp.Delivery)

var (
	deviceDbCache         sync.Map
	gatewayDbCache        sync.Map
	dataRateDbCache       sync.Map
	codingRateDbCache     sync.Map
	frequencyDbCache      sync.Map
	accuracySourceDbCache sync.Map
	userAgentDbCache      sync.Map
	userIdDbCache         sync.Map
	experimentNameDbCache sync.Map
)

type Configuration struct {
	AmqpHost     string `env:"AMQP_HOST"`
	AmqpPort     string `env:"AMQP_PORT"`
	AmqpUser     string `env:"AMQP_USER"`
	AmqpPassword string `env:"AMQP_PASSWORD"`
	AmqpExchange string `env:"AMQP_EXHANGE"`
	AmqpQueue    string `env:"AMQP_QUEUE"`

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
	AmqpHost:     "localhost",
	AmqpPort:     "5672",
	AmqpUser:     "user",
	AmqpPassword: "password",
	AmqpExchange: "new_packets",
	AmqpQueue:    "postgres_insert_raw",

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

	insertDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "ttnmapper_postgres_inserts_raw_duration",
		Help:    "How long the processing and insert of a packet takes",
		Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1, 1.5, 2, 5, 10, 100, 1000, 10000},
	})
)

func main() {

	err := gonfig.GetConf("conf.json", &myConfiguration)
	if err != nil {
		log.Println(err)
	}

	log.Printf("[Configuration]\n%s\n", prettyPrint(myConfiguration)) // output: [UserA, UserB]

	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe("0.0.0.0:"+myConfiguration.PrometheusPort, nil)
		if err != nil {
			log.Print(err.Error())
		}
	}()

	// Table name prefixes
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//return "ttnmapper_" + defaultTableName
		return defaultTableName
	}

	db, err := gorm.Open("postgres", "host="+myConfiguration.PostgresHost+" port="+myConfiguration.PostgresPort+" user="+myConfiguration.PostgresUser+" dbname="+myConfiguration.PostgresDatabase+" password="+myConfiguration.PostgresPassword+"")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if myConfiguration.PostgresDebugLog {
		db.LogMode(true)
	}

	// Create tables if they do not exist
	log.Println("Performing auto migrate")
	db.AutoMigrate(
		&types.Packet{},
		&types.Device{},
		&types.Frequency{},
		&types.DataRate{},
		&types.CodingRate{},
		&types.AccuracySource{},
		&types.Experiment{},
		&types.User{},
		&types.UserAgent{},
		&types.Gateway{},
		&types.FineTimestampKeyID{},
	)

	// Start threads to handle Postgres inserts
	log.Println("Starting database insert threads")
	for i := 0; i < myConfiguration.PostgresInsertThreads; i++ {
		go insertToPostgres(i+1, db)
	}

	// Start amqp listener on this thread - blocking function
	log.Println("Starting AMQP thread")
	subscribeToRabbit()
}

func subscribeToRabbit() {
	conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		myConfiguration.AmqpExchange, // name
		"fanout",                     // type
		true,                         // durable
		false,                        // auto-deleted
		false,                        // internal
		false,                        // no-wait
		nil,                          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		myConfiguration.AmqpQueue, // name
		true,                      // durable
		false,                     // delete when unused
		false,                     // exclusive
		false,                     // no-wait
		nil,                       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		10,    // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set queue QoS")

	err = ch.QueueBind(
		q.Name,                       // queue name
		"",                           // routing key
		myConfiguration.AmqpExchange, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	// Start thread that listens for new amqp messages
	go func() {
		for d := range msgs {
			log.Print(" [a] Packet received")
			messageChannel <- d
		}
	}()

	log.Printf("Init Complete")
	forever := make(chan bool)
	<-forever
}

func insertToPostgres(thread int, db *gorm.DB) {
	// Wait for a message and insert it into Postgres
	for d := range messageChannel {
		log.Printf("[%d][p] Processing packet", thread)

		// The message form amqp is a json string. Unmarshal to ttnmapper uplink struct
		var message types.TtnMapperUplinkMessage
		if err := json.Unmarshal(d.Body, &message); err != nil {
			log.Print("[%d][p] "+err.Error(), thread)
			continue
		}

		// Iterate gateways. We store it flat in the database
		for _, gateway := range message.Gateways {
			gatewayStart := time.Now()

			// Copy required fields in correct format into a database row struct
			entry, err := messageToEntry(db, message, gateway)
			if err != nil {
				log.Printf(err.Error())
				continue
			}

			// Insert into database
			err = db.Create(&entry).Error
			if err == nil {
				log.Printf("[%d][p] Inserted entry", thread)
				dbInserts.Inc()
			} else {
				log.Println(prettyPrint(entry))
				log.Print("[%d][p] PG Insert", thread)
				failOnError(err, "")
			}

			// Prometheus stats
			gatewayElapsed := time.Since(gatewayStart)
			insertDuration.Observe(float64(gatewayElapsed.Nanoseconds()) / 1000.0 / 1000.0) //nanoseconds to milliseconds
		}

		// If we get here all inserts were successful. Otherwise we would have quit.
		d.Ack(false)

	}
}

func messageToEntry(db *gorm.DB, message types.TtnMapperUplinkMessage, gateway types.TtnMapperGateway) (types.Packet, error) {
	var entry = types.Packet{}

	// Time
	seconds := message.Time / 1000000000
	nanos := message.Time % 1000000000
	entry.Time = time.Unix(seconds, nanos)

	// DeviceID
	deviceIndexer := types.DeviceIndexer{AppId: message.AppID, DevId: message.DevID, DevEui: message.DevEui}
	i, ok := deviceDbCache.Load(deviceIndexer)
	if ok {
		entry.DeviceID = i.(uint)
	} else {
		deviceDb := types.Device{AppId: message.AppID, DevId: message.DevID, DevEui: message.DevEui}
		err := db.FirstOrCreate(&deviceDb, &deviceDb).Error
		if err != nil {
			return entry, err
		}
		entry.DeviceID = deviceDb.ID
		deviceDbCache.Store(deviceIndexer, deviceDb.ID)
	}

	// FPort, FCnt
	entry.FPort = message.FPort
	entry.FCnt = uint32(message.FCnt)

	// FrequencyID
	i, ok = frequencyDbCache.Load(message.Frequency)
	if ok {
		entry.FrequencyID = i.(uint)
	} else {
		frequencyDb := types.Frequency{Herz: message.Frequency}
		err := db.FirstOrCreate(&frequencyDb, &frequencyDb).Error
		if err != nil {
			return entry, err
		}
		entry.FrequencyID = frequencyDb.ID
		frequencyDbCache.Store(message.Frequency, frequencyDb.ID)
	}

	// DataRateID
	dataRateIndexer := types.DataRateIndexer{
		Modulation:      message.Modulation,
		Bandwidth:       message.Bandwidth,
		SpreadingFactor: message.SpreadingFactor,
		Bitrate:         message.Bitrate}
	i, ok = dataRateDbCache.Load(dataRateIndexer)
	if ok {
		entry.DataRateID = i.(uint)
	} else {
		dataRateDb := types.DataRate{
			Modulation:      message.Modulation,
			Bandwidth:       message.Bandwidth,
			SpreadingFactor: message.SpreadingFactor,
			Bitrate:         message.Bitrate}
		err := db.FirstOrCreate(&dataRateDb, &dataRateDb).Error
		if err != nil {
			return entry, err
		}
		entry.DataRateID = dataRateDb.ID
		dataRateDbCache.Store(dataRateIndexer, dataRateDb.ID)
	}

	// CodingRateID
	i, ok = codingRateDbCache.Load(message.CodingRate)
	if ok {
		entry.CodingRateID = i.(uint)
	} else {
		codingRateDb := types.CodingRate{Name: message.CodingRate}
		err := db.FirstOrCreate(&codingRateDb, &codingRateDb).Error
		if err != nil {
			return entry, err
		}
		entry.CodingRateID = codingRateDb.ID
		codingRateDbCache.Store(message.CodingRate, codingRateDb.ID)
	}

	// GatewayID
	i, ok = gatewayDbCache.Load(gateway.GatewayId)
	if ok {
		entry.GatewayID = i.(uint)
	} else {
		gatewayDb := types.Gateway{GtwId: gateway.GatewayId, GtwEui: gateway.GatewayEui}
		err := db.FirstOrCreate(&gatewayDb, &gatewayDb).Error
		if err != nil {
			return entry, err
		}
		entry.GatewayID = gatewayDb.ID
		gatewayDbCache.Store(gateway.GatewayId, gatewayDb.ID)
	}

	// TODO update Gateway location and last heard - use a different microservice for this

	// GatewayTime
	if gateway.Time != 0 {
		seconds = gateway.Time / 1000000000
		nanos = gateway.Time % 1000000000
		gatewayTime := time.Unix(seconds, nanos)
		entry.GatewayTime = &gatewayTime
	}

	// Timestamp
	if gateway.Timestamp != 0 {
		entry.Timestamp = &gateway.Timestamp
	}

	// FineTimestamp
	if gateway.FineTimestamp != 0 {
		entry.FineTimestamp = &gateway.FineTimestamp
	}

	// FineTimestampEncrypted
	if len(gateway.FineTimestampEncrypted) > 0 {
		entry.FineTimestampEncrypted = &gateway.FineTimestampEncrypted
	}

	// FineTimestampKeyID
	if gateway.FineTimestampEncryptedKeyId != "" {
		// TODO: cache if this is done often
		fineTimestampKeyId := types.FineTimestampKeyID{FineTimestampEncryptedKeyId: gateway.FineTimestampEncryptedKeyId}
		err := db.FirstOrCreate(&fineTimestampKeyId, &fineTimestampKeyId).Error
		if err != nil {
			return entry, err
		}
		entry.FineTimestampKeyID = &fineTimestampKeyId.ID
	}

	// ChannelIndex
	entry.ChannelIndex = gateway.ChannelIndex

	// Rssi, SignalRssi, Snr
	entry.Rssi = gateway.Rssi
	if gateway.SignalRssi != 0 {
		entry.SignalRssi = &gateway.SignalRssi
	}
	if gateway.Snr != 0 {
		entry.Snr = &gateway.Snr
	}

	// Latitude, Longitude, Altitude, AccuracyMeters, Satellites, Hdop
	entry.Latitude = message.Latitude
	entry.Longitude = message.Longitude
	entry.Altitude = message.Altitude
	if message.AccuracyMeters != 0 {
		entry.AccuracyMeters = &message.AccuracyMeters
	}
	if message.Satellites != 0 {
		entry.Satellites = &message.Satellites
	}
	if message.Hdop != 0 {
		hdop := math.Min(message.Hdop, 999.9) // database field is 3.1 - cap otherwise we get an error
		entry.Hdop = &hdop
	}

	// AccuracySourceID
	i, ok = accuracySourceDbCache.Load(message.AccuracySource)
	if ok {
		entry.AccuracySourceID = i.(uint)
	} else {
		accuracySourceDb := types.AccuracySource{Name: message.AccuracySource}
		err := db.FirstOrCreate(&accuracySourceDb, &accuracySourceDb).Error
		if err != nil {
			return entry, err
		}
		entry.AccuracySourceID = accuracySourceDb.ID
		accuracySourceDbCache.Store(message.AccuracySource, accuracySourceDb.ID)
	}

	// ExperimentID
	if message.Experiment != "" {
		i, ok = experimentNameDbCache.Load(message.Experiment)
		if ok {
			experimentId := i.(uint)
			entry.ExperimentID = &experimentId
		} else {
			experimentNameDb := types.Experiment{Name: message.Experiment}
			err := db.FirstOrCreate(&experimentNameDb, &experimentNameDb).Error
			if err != nil {
				return entry, err
			}
			entry.ExperimentID = &experimentNameDb.ID
			experimentNameDbCache.Store(message.Experiment, experimentNameDb.ID)
		}
	}

	// UserID
	i, ok = userIdDbCache.Load(message.UserId)
	if ok {
		entry.UserID = i.(uint)
	} else {
		userIdDb := types.User{Identifier: message.UserId}
		err := db.FirstOrCreate(&userIdDb, &userIdDb).Error
		if err != nil {
			return entry, err
		}
		entry.UserID = userIdDb.ID
		userIdDbCache.Store(message.UserId, userIdDb.ID)
	}

	// UserAgentID
	i, ok = userAgentDbCache.Load(message.UserAgent)
	if ok {
		entry.UserAgentID = i.(uint)
	} else {
		userAgentDb := types.UserAgent{Name: message.UserAgent}
		err := db.FirstOrCreate(&userAgentDb, &userAgentDb).Error
		if err != nil {
			return entry, err
		}
		entry.UserAgentID = userAgentDb.ID
		userAgentDbCache.Store(message.UserAgent, userAgentDb.ID)
	}

	return entry, nil
}
