package database

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	// Caches for static values, do not expire
	deviceDbCache         sync.Map
	antennaDbCache        sync.Map
	dataRateDbCache       sync.Map
	codingRateDbCache     sync.Map
	frequencyDbCache      sync.Map
	accuracySourceDbCache sync.Map
	userAgentDbCache      sync.Map
	userIdDbCache         sync.Map
	experimentNameDbCache sync.Map

	// Caches for dynamic content, needs to expire
	gatewayDbCache             *cache.Cache
	networkOnlineGatewaysCache *cache.Cache

	Db *gorm.DB
)

type Context struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
	DebugLog bool
}

func (databaseContext *Context) Init() {

	var gormLogLevel = logger.Silent
	if databaseContext.DebugLog {
		log.Println("Database debug logging enabled")
		gormLogLevel = logger.Info
	}

	// Postgres has a max length for application name of 63 chars. If we do not limit this we get a "received unexpected message" error
	applicationName := filepath.Base(os.Args[0])
	if len(applicationName) > 63 {
		applicationName = applicationName[:63]
	}

	dsn := "host=" + databaseContext.Host + " port=" + databaseContext.Port + " user=" + databaseContext.User +
		" dbname=" + databaseContext.Database + " password=" + databaseContext.Password + " sslmode=disable" +
		" application_name=" + applicationName
	log.Println(dsn)
	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:          logger.Default.LogMode(gormLogLevel),
		CreateBatchSize: 1000,
	})
	if err != nil {
		panic(err.Error())
	}

	// Init caches
	gatewayDbCache = cache.New(5*time.Minute, 1*time.Minute)
	networkOnlineGatewaysCache = cache.New(5*time.Minute, 1*time.Minute)
}

func AutoMigrate(models ...interface{}) {
	// Create tables if they do not exist
	log.Println("Performing auto migrate")
	if err := Db.AutoMigrate(
		//	&Device{},
		//	&Frequency{},
		//	&DataRate{},
		//	&CodingRate{},
		//	&AccuracySource{},
		//	&Experiment{},
		//	&User{},
		//	&UserAgent{},
		//	&Antenna{},
		//	&FineTimestampKeyID{},
		//	&Packet{},
		models...,
	); err != nil {
		log.Println("Unable autoMigrateDB - " + err.Error())
	}
}

func GetAllTtsNetworksToFetch() []TtsV3FetchStatus {
	var networks []TtsV3FetchStatus
	Db.Find(&networks)
	return networks
}
