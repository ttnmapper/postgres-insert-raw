package database

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"testing"
	"time"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

type Configuration struct {
	PostgresHost     string `envconfig:"POSTGRES_HOST"`
	PostgresPort     string `envconfig:"POSTGRES_PORT"`
	PostgresUser     string `envconfig:"POSTGRES_USER"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD"`
	PostgresDatabase string `envconfig:"POSTGRES_DATABASE"`
	PostgresDebugLog bool   `envconfig:"POSTGRES_DEBUG_LOG"`
}

var myConfiguration = Configuration{
	PostgresHost:     "localhost",
	PostgresPort:     "5432",
	PostgresUser:     "username",
	PostgresPassword: "password",
	PostgresDatabase: "database",
	PostgresDebugLog: true,
}

func initDb() {
	err := envconfig.Process("", &myConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Init database")
	databaseContext := Context{
		Host:     myConfiguration.PostgresHost,
		Port:     myConfiguration.PostgresPort,
		User:     myConfiguration.PostgresUser,
		Database: myConfiguration.PostgresDatabase,
		Password: myConfiguration.PostgresPassword,
		DebugLog: myConfiguration.PostgresDebugLog,
	}
	databaseContext.Init()
}

func TestGetGatewaysByNameOrId(t *testing.T) {
	initDb()

	gateways := GetGatewaysByNameOrId("Xf1p4e7QPx2Kb1lqLq4ql/H9zF+2jH0lLpMGg+w4W7o=")
	log.Println(utils.PrettyPrint(gateways))
}

func TestGetPacketsForGateway(t *testing.T) {
	initDb()

	packets, err := GetPacketsForGateway("NS_TTS_V3://deutschebahn@000013", "Xf1p4e7QPx2Kb1lqLq4ql/H9zF+2jH0lLpMGg+w4W7o=", time.Unix(0, 0), time.Now(), 10000)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// This should return data, but it does not
	log.Println(utils.PrettyPrint(packets))
}
