package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/database"
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
	//initTests()
	FetchPbRoutingPolicies()
}
