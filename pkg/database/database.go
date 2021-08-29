package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
	"ttnmapper-postgres-insert-raw/pkg/types"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

var (
	deviceDbCache         sync.Map
	antennaDbCache        sync.Map
	dataRateDbCache       sync.Map
	codingRateDbCache     sync.Map
	frequencyDbCache      sync.Map
	accuracySourceDbCache sync.Map
	userAgentDbCache      sync.Map
	userIdDbCache         sync.Map
	experimentNameDbCache sync.Map
	db                    *gorm.DB
)

type DatabaseContext struct {
	Host     string
	Port     string
	User     string
	Database string
	Password string
	DebugLog bool
}

func (databaseContext *DatabaseContext) Init() {

	var gormLogLevel = logger.Silent
	if databaseContext.DebugLog {
		log.Println("Database debug logging enabled")
		gormLogLevel = logger.Info
	}

	dsn := "host=" + databaseContext.Host + " port=" + databaseContext.Port + " user=" + databaseContext.User + " dbname=" + databaseContext.Database + " password=" + databaseContext.Password + " sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		panic(err.Error())
	}

	// Create tables if they do not exist
	//log.Println("Performing auto migrate")
	//if err := db.AutoMigrate(
	//	&types.Packet{},
	//	&types.Device{},
	//	&types.Frequency{},
	//	&types.DataRate{},
	//	&types.CodingRate{},
	//	&types.AccuracySource{},
	//	&types.Experiment{},
	//	&types.User{},
	//	&types.UserAgent{},
	//	&types.Antenna{},
	//	&types.FineTimestampKeyID{},
	//); err != nil {
	//	log.Println("Unable autoMigrateDB - " + err.Error())
	//}
}

func InsertEntry(entry *types.Packet) error {
	err := db.Create(&entry).Error
	return err
}

func UplinkMessageToPacket(message types.TtnMapperUplinkMessage, gateway types.TtnMapperGateway) (types.Packet, error) {
	var entry = types.Packet{}

	// Time
	seconds := message.Time / 1000000000
	nanos := message.Time % 1000000000
	entry.Time = time.Unix(seconds, nanos)

	// DeviceID
	deviceIndexer := types.DeviceIndexer{NetworkId: message.NetworkId, AppId: message.AppID, DevId: message.DevID, DevEui: message.DevEui}
	i, ok := deviceDbCache.Load(deviceIndexer)
	if ok {
		entry.DeviceID = i.(uint)
	} else {
		log.Println("Get or create device from/in DB:", deviceIndexer)
		deviceDb := types.Device{NetworkId: message.NetworkId, AppId: message.AppID, DevId: message.DevID, DevEui: message.DevEui}
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

	// AntennaID - packets are stored with a pointer to the antenna that received it. A network has multiple gateways, a gateway has multiple antennas.
	// We therefore store coverage data per antenna, assuming antenna index 0 when we don't know the antenna index.
	antennaIndexer := types.AntennaIndexer{NetworkId: gateway.NetworkId, GatewayId: gateway.GatewayId, AntennaIndex: gateway.AntennaIndex}
	i, ok = antennaDbCache.Load(antennaIndexer)
	if ok {
		entry.AntennaID = i.(uint)
	} else {
		antennaDb := types.Antenna{NetworkId: gateway.NetworkId, GatewayId: gateway.GatewayId, AntennaIndex: gateway.AntennaIndex}
		err := db.FirstOrCreate(&antennaDb, &antennaDb).Error
		if err != nil {
			return entry, err
		}
		entry.AntennaID = antennaDb.ID
		antennaDbCache.Store(antennaIndexer, antennaDb.ID)
	}

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
	entry.Snr = gateway.Snr

	// Latitude, Longitude, Altitude, AccuracyMeters, Satellites, Hdop
	entry.Latitude = utils.CapFloatTo(message.Latitude, 10, 6)
	entry.Longitude = utils.CapFloatTo(message.Longitude, 10, 6)
	entry.Altitude = utils.CapFloatTo(message.Altitude, 6, 1)

	if message.AccuracyMeters != 0 {
		accuracy := utils.CapFloatTo(message.AccuracyMeters, 6, 2)
		entry.AccuracyMeters = &accuracy
	}
	if message.Satellites != 0 {
		entry.Satellites = &message.Satellites
	}
	if message.Hdop != 0 {
		hdop := utils.CapFloatTo(message.Hdop, 3, 1)
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