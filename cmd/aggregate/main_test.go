package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func IniDb() {
	err := envconfig.Process("", &myConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}

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

func TestMigrateDb(t *testing.T) {
	IniDb()

	database.AutoMigrate(&database.GridCell{}, &database.RadarBeam{})
}
