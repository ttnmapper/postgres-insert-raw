package tests

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/types"
)

type Configuration struct {
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`
	PostgresDebugLog bool   `env:"POSTGRES_DEBUG_LOG"`
}

var myConfiguration = Configuration{
	PostgresHost:     "localhost",
	PostgresPort:     "5432",
	PostgresUser:     "username",
	PostgresPassword: "password",
	PostgresDatabase: "database",
	PostgresDebugLog: true,
}

func TestMigrateDb(t *testing.T) {

	myConfiguration.PostgresHost = os.Getenv("POSTGRES_HOST")
	myConfiguration.PostgresPort = os.Getenv("POSTGRES_PORT")
	myConfiguration.PostgresUser = os.Getenv("POSTGRES_USER")
	myConfiguration.PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
	myConfiguration.PostgresDatabase = os.Getenv("POSTGRES_DATABASE")

	log.Printf("[Configuration]\n%v\n", myConfiguration)

	dsn := "host=" + myConfiguration.PostgresHost + " port=" + myConfiguration.PostgresPort + " user=" + myConfiguration.PostgresUser + " dbname=" + myConfiguration.PostgresDatabase + " password=" + myConfiguration.PostgresPassword + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

	log.Println("Performing auto migrate")
	if err := db.AutoMigrate(
		&types.Packet{},
		&types.Device{},
		&types.Frequency{},
		&types.DataRate{},
		&types.CodingRate{},
		&types.AccuracySource{},
		&types.Experiment{},
		&types.User{},
		&types.UserAgent{},
		&types.Antenna{},
		&types.FineTimestampKeyID{},
	); err != nil {
		log.Println("Unable autoMigrateDB - ", err.Error())
	}
}
