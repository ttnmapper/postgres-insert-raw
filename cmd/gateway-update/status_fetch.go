package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
	"ttnmapper-postgres-insert-raw/pkg/database"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/helium"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/noc"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/packet_broker"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/thethingsstack"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/web"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

func subscribeToRabbitRaw() {
	// Start thread that listens for new amqp messages
	go func() {
		conn, err := amqp.Dial("amqp://" + myConfiguration.AmqpUser + ":" + myConfiguration.AmqpPassword + "@" + myConfiguration.AmqpHost + ":" + myConfiguration.AmqpPort + "/")
		utils.FailOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		// Create a channel for errors
		notify := conn.NotifyClose(make(chan *amqp.Error)) //error channel

		ch, err := conn.Channel()
		utils.FailOnError(err, "Failed to open a channel")
		defer ch.Close()

		err = ch.ExchangeDeclare(
			myConfiguration.AmqpExchangeRawPackets, // name
			"fanout",                               // type
			true,                                   // durable
			false,                                  // auto-deleted
			false,                                  // internal
			false,                                  // no-wait
			nil,                                    // arguments
		)
		utils.FailOnError(err, "Failed to declare an exchange")

		q, err := ch.QueueDeclare(
			myConfiguration.AmqpQueueRawPackets, // name
			false,                               // durable
			false,                               // delete when unused
			false,                               // exclusive
			false,                               // no-wait
			nil,                                 // arguments
		)
		utils.FailOnError(err, "Failed to declare a queue")

		err = ch.Qos(
			10,    // prefetch count
			0,     // prefetch size
			false, // global
		)
		utils.FailOnError(err, "Failed to set queue QoS")

		err = ch.QueueBind(
			q.Name,                                 // queue name
			"",                                     // routing key
			myConfiguration.AmqpExchangeRawPackets, // exchange
			false,
			nil)
		utils.FailOnError(err, "Failed to bind a queue")

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		utils.FailOnError(err, "Failed to register a consumer")

		log.Println("AMQP started")

	waitForMessages:
		for {
			select {
			case err := <-notify:
				if err != nil {
					log.Println(err.Error())
				}
				break waitForMessages
			case d := <-msgs:
				//log.Printf(" [a] Packet received")
				rawPacketsChannel <- d
			}
		}

		log.Fatal("Subscribe channel closed")
	}()

	// Start the thread that processes new amqp messages
	go func() {
		processRawPackets()
	}()
}

func startPeriodicFetchers() {
	//nocTicker := time.NewTicker(time.Duration(myConfiguration.FetchNocInterval) * time.Second)
	webTicker := time.NewTicker(time.Duration(myConfiguration.FetchWebInterval) * time.Second)
	pbTicker := time.NewTicker(time.Duration(myConfiguration.FetchPacketBrokerInterval) * time.Second)
	heliumTicker := time.NewTicker(time.Duration(myConfiguration.FetchHeliumInterval) * time.Second)
	ttsTicker := time.NewTicker(time.Duration(myConfiguration.FetchTtsInterval) * time.Second)
	pbRoutingTicker := time.NewTicker(time.Duration(myConfiguration.FetchRoutingInterval) * time.Second)

	go func() {
		for {
			select {
			//case <-nocTicker.C:
			//	if myConfiguration.FetchNoc {
			//		go fetchNocStatuses()
			//	}
			case <-webTicker.C:
				if myConfiguration.FetchWeb {
					go fetchWebStatuses()
				}
			case <-pbTicker.C:
				if myConfiguration.FetchPacketBroker {
					go fetchPacketBrokerStatuses()
				}
			case <-heliumTicker.C:
				if myConfiguration.FetchHelium {
					go fetchHeliumStatuses()
				}
			case <-ttsTicker.C:
				if myConfiguration.FetchTts {
					go fetchTtsStatuses()
				}
			case <-pbRoutingTicker.C:
				if myConfiguration.FetchRouting {
					go FetchPbRoutingPolicies()
				}
			}
		}
	}()
}

var busyFetchingNoc = false

func fetchNocStatuses() {
	if busyFetchingNoc {
		return
	}
	busyFetchingNoc = true

	gateways, err := noc.FetchNocStatuses()
	if err != nil {
		log.Println(err.Error())
	} else {
		for id, gateway := range gateways {
			ttnMapperGateway := noc.NocGatewayToTtnMapperGateway(id, gateway)
			log.Print("NOC ", "", "\t", ttnMapperGateway.GatewayId+"\t", ttnMapperGateway.Time)
			UpdateGateway(ttnMapperGateway)
		}
	}

	busyFetchingNoc = false
}

var busyFetchingWeb = false

func fetchWebStatuses() {
	if busyFetchingWeb {
		return
	}
	busyFetchingWeb = true

	gatewayCount := 0
	gateways, err := web.FetchWebStatuses()
	if err != nil {
		log.Println(err.Error())
	} else {
		for _, gateway := range gateways {
			gatewayCount++
			ttnMapperGateway := web.WebGatewayToTtnMapperGateway(*gateway)
			log.Print("WEB ", "", "\t", ttnMapperGateway.GatewayId+"\t", ttnMapperGateway.Time)
			UpdateGateway(ttnMapperGateway)
		}
	}

	log.Printf("Fetched %d gateways from TTN website", gatewayCount)
	busyFetchingWeb = false
}

/*
Fetching on 2021-11-17 took
real	4m5.426s
user	0m3.360s
sys	0m4.620s
*/
var busyFetchingPacketBroker = false

func fetchPacketBrokerStatuses() {
	if busyFetchingPacketBroker {
		return
	}
	busyFetchingPacketBroker = true

	gatewayCount := 0
	page := 0
	for {

		gateways, err := packet_broker.FetchStatuses(page)
		if err != nil {
			log.Println(err.Error())
			break
		} else {
			for _, gateway := range gateways {
				gatewayCount++
				ttnMapperGateway, err := packet_broker.PbGatewayToTtnMapperGateway(gateway)
				if err == nil {
					log.Print("PB ", "", "\t", ttnMapperGateway.GatewayId+"\t", ttnMapperGateway.Time)
					UpdateGateway(ttnMapperGateway)
				}
			}
		}
		page++
	}

	log.Printf("Fetched %d gateways from Packet Broker", gatewayCount)
	busyFetchingPacketBroker = false
}

/*
Fetching Helium on 2021-11-17 took
real    306m29.503s
user    1m47.352s
sys     3m13.356s
*/
var busyFetchingHelium = false

func fetchHeliumStatuses() {
	if busyFetchingHelium {
		return
	}
	busyFetchingHelium = true

	hotspotCount := 0
	cursor := ""
	for {
		response, err := helium.FetchStatuses(cursor)
		if err != nil {
			log.Println(err.Error())
			break
		}

		log.Printf("HELIUM %d hotspots\n", len(response.Data))

		for _, hotspot := range response.Data {
			hotspotCount++
			ttnMapperGateway, err := helium.HeliumHotspotToTtnMapperGateway(hotspot)
			if err == nil {
				//log.Print("HELIUM ", "", "\t", ttnMapperGateway.GatewayId+"\t", ttnMapperGateway.Time)
				UpdateGateway(ttnMapperGateway)
			}
		}

		cursor = response.Cursor
		if cursor == "" {
			log.Println("Cursor empty")
			break
		}
		if len(response.Data) == 0 {
			log.Println("No hotspots in response")
			break
		}
	}

	log.Printf("Fetched %d hotspots from Helium", hotspotCount)
	busyFetchingHelium = false
}

// Fetch statuses from private TTS instances

var busyFetchingTtsNetworkStatuses = false

func fetchTtsStatuses() {
	if busyFetchingTtsNetworkStatuses {
		return
	}
	busyFetchingTtsNetworkStatuses = true

	networks := database.GetAllTtsNetworksToFetch()

	gatewayCount := 0
	for _, network := range networks {
		gateways, err := thethingsstack.FetchGateways(network.TenantId, network.ApiKey)
		if err != nil {
			log.Println(err.Error())
		}

		// Fetch gateway statuses in batches
		batchSize := 50
		for i := 0; i < len(gateways); i += batchSize {
			log.Printf("[TTS API] Fetching batch of %d gateway statuses", batchSize)
			endIndex := i + batchSize
			if len(gateways) < endIndex {
				endIndex = len(gateways)
			}
			currentlyFetchingGateways := gateways[i:endIndex]
			gatewayStatuses, err := thethingsstack.FetchStatusesBatch(currentlyFetchingGateways, network.ApiKey)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			if gatewayStatuses.Entries == nil {
				log.Println("[TTS API] Status Entries is nil")
			}
			// Iterate status responses
			for gatewayId, status := range gatewayStatuses.Entries {
				// Iterate fetched gateway list to find requested gateway's data
				for _, gateway := range currentlyFetchingGateways {
					// If we found the gateway, ie the id matches, update its status
					if gateway.Ids.GatewayId == gatewayId {
						log.Println(network.TenantId, gatewayId)
						ttnMapperGateway, err := thethingsstack.TtsApiGatewayToTtnMapperGateway(network.TenantId, gateway, status)
						if err != nil {
							log.Println(err)
							continue
						}
						UpdateGateway(ttnMapperGateway)
						gatewayCount++
					}
				}
			}
		}
		log.Printf("[TTS API] Fetched %d gateway statuses for network %s", gatewayCount, network.TenantId)
	}

	busyFetchingTtsNetworkStatuses = false
}

// Fetch routing policies from Packet Broker

var busyFetchingRouting = false

func FetchPbRoutingPolicies() {
	if busyFetchingRouting {
		return
	}
	busyFetchingRouting = true

	var netId uint32 = 0x000013
	tenantId := "ttn"
	policies := packet_broker.FetchRoutingPolicies(netId, tenantId, os.Getenv("PB_API_KEY_ID"), os.Getenv("PB_API_KEY_SECRET"))

	//log.Println(utils.PrettyPrint(policies))
	for _, policy := range policies {
		// The results include wildcard policies. Replace empty wildcard fields with known network values.
		policy.HomeNetworkNetId = netId
		policy.HomeNetworkTenantId = tenantId

		dbPolicy := database.PacketBrokerRoutingPolicy{}
		packet_broker.RoutingPolicyToDbPolicy(policy, &dbPolicy)

		//log.Println(dbPolicy.HomeNetworkId, " - ", dbPolicy.ForwarderNetworkId)
		database.InsertOrUpdateRoutingPolicy(dbPolicy)
	}

	busyFetchingRouting = false
}
