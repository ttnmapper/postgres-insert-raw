package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"strings"
	"sync"
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
	//); err != nil {
	//	log.Println("Unable autoMigrateDB - " + err.Error())
	//}
}

func InsertEntry(entry *Packet) error {
	err := db.Create(&entry).Error
	return err
}

func InsertGatewayLocationsBatch(gatewayLocations []GatewayLocation) error {
	tx := db.Begin()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, location := range gatewayLocations {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, location.NetworkId)
		valueArgs = append(valueArgs, location.GatewayId)
		valueArgs = append(valueArgs, location.InstalledAt)
		valueArgs = append(valueArgs, location.Latitude)
		valueArgs = append(valueArgs, location.Longitude)
		valueArgs = append(valueArgs, location.Altitude)
	}

	stmt := fmt.Sprintf("INSERT INTO gateway_locations (network_id, gateway_id, installed_at, latitude, longitude, altitude) VALUES %s", strings.Join(valueStrings, ","))
	err := tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	//tx.Rollback()
	return err
}

func InsertPacketsBatch(packets []Packet) error {
	tx := db.Begin()
	valueStrings := []string{}
	valueArgs := []interface{}{}
	fieldNames := "time, device_id, f_port, f_cnt, frequency_id, data_rate_id, coding_rate_id, " +
		"antenna_id, gateway_time, timestamp, " +
		"fine_timestamp, fine_timestamp_encrypted, fine_timestamp_key_id, " +
		"channel_index, rssi, signal_rssi, snr, " +
		"latitude, longitude, altitude, accuracy_meters, satellites, hdop, accuracy_source_id, " +
		"experiment_id, user_id, user_agent_id, deleted_at"

	// ('2016-01-31 15:50:01', 185203, 0, 0, 123, 1, 1, 32699, NULL, NULL, NULL, NULL, NULL, 0, -109.000000, NULL, 2.500000, 52.244205, 6.856759, 0.000000, NULL, NULL, NULL, 1, NULL, 1, 1, NULL),
	// ('2016-01-31 15:50:01', 185203, 0, 0, 123, 1, 1, 160, NULL, NULL, NULL, NULL, NULL, 0, -119.000000, NULL, -4.800000, 52.244205, 6.856759, 0.000000, NULL, NULL, NULL, 1, NULL, 1, 1, NULL),
	// ('2016-01-31 15:49:56', 185203, 0, 0, 1, 1, 1, 32699, NULL, NULL, NULL, NULL, NULL, 0, -107.000000, NULL, 7.000000, 52.243984, 6.856919, 0.000000, NULL, NULL, NULL, 1, NULL, 1, 1, NULL),
	// ('2016-01-31 15:49:56', 185203, 0, 0, 1, 1, 1, 160, NULL, NULL, NULL, NULL, NULL, 0, -113.000000, NULL, 0.500000, 52.243984, 6.856919, 0.000000, NULL, NULL, NULL, 1, NULL, 1, 1, NULL)
	for _, packet := range packets {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, packet.Time)
		valueArgs = append(valueArgs, packet.DeviceID)
		valueArgs = append(valueArgs, packet.FPort)
		valueArgs = append(valueArgs, packet.FCnt)
		valueArgs = append(valueArgs, packet.FrequencyID)
		valueArgs = append(valueArgs, packet.DataRateID)
		valueArgs = append(valueArgs, packet.CodingRateID)
		valueArgs = append(valueArgs, packet.AntennaID)
		valueArgs = append(valueArgs, packet.GatewayTime)
		valueArgs = append(valueArgs, packet.Timestamp)
		valueArgs = append(valueArgs, packet.FineTimestamp)
		valueArgs = append(valueArgs, packet.FineTimestampEncrypted)
		valueArgs = append(valueArgs, packet.FineTimestampKeyID)
		valueArgs = append(valueArgs, packet.ChannelIndex)
		valueArgs = append(valueArgs, packet.Rssi)
		valueArgs = append(valueArgs, packet.SignalRssi)
		valueArgs = append(valueArgs, packet.Snr)
		valueArgs = append(valueArgs, packet.Latitude)
		valueArgs = append(valueArgs, packet.Longitude)
		valueArgs = append(valueArgs, packet.Altitude)
		valueArgs = append(valueArgs, packet.AccuracyMeters)
		valueArgs = append(valueArgs, packet.Satellites)
		valueArgs = append(valueArgs, packet.Hdop)
		valueArgs = append(valueArgs, packet.AccuracySourceID)
		valueArgs = append(valueArgs, packet.ExperimentID)
		valueArgs = append(valueArgs, packet.UserID)
		valueArgs = append(valueArgs, packet.UserAgentID)
		valueArgs = append(valueArgs, packet.DeletedAt)
	}

	stmt := fmt.Sprintf("INSERT INTO packets (%s) VALUES %s", fieldNames, strings.Join(valueStrings, ","))
	err := tx.Exec(stmt, valueArgs...).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	//tx.Rollback()
	return err
}

func GetAllTtnV2Antennas() []Antenna {
	var antennas []Antenna
	db.Where("network_id LIKE 'NS_TTN_V2://%'").Find(&antennas)
	return antennas
}

func FindAntenna(networkId string, gatewayId string, antennaIndex uint8) Antenna {
	antenna := Antenna{NetworkId: networkId, GatewayId: gatewayId, AntennaIndex: antennaIndex}
	db.FirstOrCreate(&antenna, &antenna)
	return antenna
}

func UpdatePacketsAntennaId(oldAntennaId uint, newAntennaId uint) {
	db.Model(&Packet{}).Where("antenna_id = ?", oldAntennaId).Update("antenna_id", newAntennaId)
}
