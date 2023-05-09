package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/database"
	"ttnmapper-postgres-insert-raw/pkg/gateway-statuses/thethingsstack"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

func initTests() {
	err := envconfig.Process("", &myConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("[Configuration]\n%s\n", utils.PrettyPrint(myConfiguration)) // output: [UserA, UserB]

	log.Println("Init database")
	databaseContext := database.Context{
		Host:     myConfiguration.PostgresHost,
		Port:     myConfiguration.PostgresPort,
		User:     myConfiguration.PostgresUser,
		Database: myConfiguration.PostgresDatabase,
		Password: myConfiguration.PostgresPassword,
		DebugLog: myConfiguration.PostgresDebugLog,
	}
	databaseContext.Init()
}

//func TestFectNocStatuses(t *testing.T) {
//	initTests()
//	fetchNocStatuses()
//}

//func TestFetchWebStatuses(t *testing.T) {
//	initTests()
//	fetchWebStatuses()
//}

func TestFetchPacketBrokerStatuses(t *testing.T) {
	initTests()
	fetchPacketBrokerStatuses()
}

func TestFetchHeliumStatuses(t *testing.T) {
	initTests()
	fetchHeliumStatuses()
}

func TestFetchHeliumSnapshot(t *testing.T) {
	initTests()
	fetchHeliumSnapshot()
}

func TestFetchTtsStatuses(t *testing.T) {
	initTests()
	fetchTtsStatuses()
}

func TestFetchPbRoutingPolicies(t *testing.T) {
	initTests()
	FetchPbRoutingPolicies()
}

func TestFetchTtsNetwork(t *testing.T) {

	tenantId := "redyteliot"
	apiKey := "NNSXS."

	gateways, err := thethingsstack.FetchGateways(tenantId, apiKey)
	if err != nil {
		log.Println(err.Error())
	}

	// Fetch gateway statuses in batches
	gatewayCount := 0
	batchSize := 50
	for i := 0; i < len(gateways); i += batchSize {
		log.Printf("[TTS API] Fetching batch of %d gateway statuses", batchSize)
		endIndex := i + batchSize
		if len(gateways) < endIndex {
			endIndex = len(gateways)
		}
		currentlyFetchingGateways := gateways[i:endIndex]
		gatewayStatuses, err := thethingsstack.FetchStatusesBatch(currentlyFetchingGateways, apiKey)
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
					log.Println(tenantId, gatewayId)
					ttnMapperGateway, err := thethingsstack.TtsApiGatewayToTtnMapperGateway(tenantId, gateway, *status)
					if err != nil {
						log.Println(err)
						continue
					}
					//UpdateGateway(ttnMapperGateway)
					log.Println(utils.PrettyPrint(ttnMapperGateway))
					gatewayCount++
				}
			}
		}
	}
	log.Printf("[TTS API] Fetched %d gateway statuses for network %s", gatewayCount, tenantId)
}
