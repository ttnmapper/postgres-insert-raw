package database

import (
	"github.com/kelseyhightower/envconfig"
	"log"
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
