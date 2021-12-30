package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	// Caches for static values, do not expire
	deviceDbCache         sync.Map
	antennaDbCache        sync.Map
	gatewayDbCache        sync.Map
	dataRateDbCache       sync.Map
	codingRateDbCache     sync.Map
	frequencyDbCache      sync.Map
	accuracySourceDbCache sync.Map
	userAgentDbCache      sync.Map
	userIdDbCache         sync.Map
	experimentNameDbCache sync.Map

	// Caches for dynamic content, needs to expire
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
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		panic(err.Error())
	}

	// Init caches
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

func InsertEntry(entry *Packet) error {
	err := Db.Create(&entry).Error
	return err
}

func InsertGatewayLocationsBatch(gatewayLocations []GatewayLocation) error {
	if len(gatewayLocations) == 0 {
		return errors.New("nothing to insert")
	}

	tx := Db.Begin()
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

func GetPacketsForAntennaAfter(antenna Antenna, afterTime time.Time) (*sql.Rows, error) {
	// Get all existing packets since gateway last moved
	return Db.Model(&Packet{}).Where("antenna_id = ? AND time > ? AND experiment_id IS NULL", antenna.ID, afterTime).Rows() // server side cursor
}

func GetPacketsForDevice(networkId string, applicationId string, deviceId string, startTime time.Time, endTime time.Time, limit int) (*sql.Rows, error) {
	session := Db.Model(&Packet{})
	session = session.Select("packets.id, packets.time, packets.f_port, packets.f_cnt, packets.gateway_time, fine_timestamp, packets.channel_index, packets.rssi, packets.signal_rssi, packets.snr, packets.latitude, packets.longitude, packets.altitude, packets.accuracy_meters, packets.satellites, packets.hdop, app_id, dev_id, dev_eui, d.network_id as device_network_id, f.herz as frequency, modulation, bandwidth, spreading_factor, bitrate, cr.name as coding_rate, a.network_id as gateway_network_id, gateway_id, antenna_index, \"as\".name as accuracy_source, ua.name as user_agent, e.name as experiment")
	session = session.Joins("JOIN devices d on packets.device_id = d.id")
	session = session.Joins("JOIN frequencies f on packets.frequency_id = f.id")
	session = session.Joins("JOIN data_rates dr on packets.data_rate_id = dr.id")
	session = session.Joins("JOIN coding_rates cr on packets.coding_rate_id = cr.id")
	session = session.Joins("JOIN antennas a on packets.antenna_id = a.id")
	session = session.Joins("JOIN accuracy_sources \"as\" on packets.accuracy_source_id = \"as\".id")
	session = session.Joins("JOIN user_agents ua on packets.user_agent_id = ua.id")
	session = session.Joins("JOIN users on packets.user_id = users.id")
	session = session.Joins("LEFT JOIN experiments e on packets.experiment_id = e.id")

	//session = session.Where("experiment_id IS NULL")
	session = session.Where("d.dev_id = ?", deviceId)
	session = session.Where("time > ? AND time < ?", startTime, endTime)
	if networkId != "" {
		session = session.Where("d.network_id = ?", networkId)
	}
	if applicationId != "" {
		session = session.Where("d.app_id = ?", applicationId)
	}

	session = session.Limit(limit)

	return session.Rows()
}

func GetPacketsForGateway(networkId string, gatewayId string, startTime time.Time, endTime time.Time, limit int) (*sql.Rows, error) {
	session := Db.Model(&Packet{})
	session = session.Select("packets.id, packets.time, packets.f_port, packets.f_cnt, packets.gateway_time, fine_timestamp, packets.channel_index, packets.rssi, packets.signal_rssi, packets.snr, packets.latitude, packets.longitude, packets.altitude, packets.accuracy_meters, packets.satellites, packets.hdop, app_id, dev_id, dev_eui, d.network_id as device_network_id, f.herz as frequency, modulation, bandwidth, spreading_factor, bitrate, cr.name as coding_rate, a.network_id as gateway_network_id, gateway_id, antenna_index, \"as\".name as accuracy_source, ua.name as user_agent, e.name as experiment")
	session = session.Joins("JOIN devices d on packets.device_id = d.id")
	session = session.Joins("JOIN frequencies f on packets.frequency_id = f.id")
	session = session.Joins("JOIN data_rates dr on packets.data_rate_id = dr.id")
	session = session.Joins("JOIN coding_rates cr on packets.coding_rate_id = cr.id")
	session = session.Joins("JOIN antennas a on packets.antenna_id = a.id")
	session = session.Joins("JOIN accuracy_sources \"as\" on packets.accuracy_source_id = \"as\".id")
	session = session.Joins("JOIN user_agents ua on packets.user_agent_id = ua.id")
	session = session.Joins("JOIN users on packets.user_id = users.id")
	session = session.Joins("LEFT JOIN experiments e on packets.experiment_id = e.id")

	//session = session.Where("experiment_id IS NULL")
	session = session.Where("a.gateway_id = ?", gatewayId)
	session = session.Where("time > ? AND time < ?", startTime, endTime)
	if networkId != "" {
		session = session.Where("a.network_id = ?", networkId)
	}

	session = session.Limit(limit)

	return session.Rows()
}

func InsertPacketsBatch(packets []Packet) error {
	if len(packets) == 0 {
		return errors.New("nothing to insert")
	}

	tx := Db.Begin()
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

func GetAllGateways() []Gateway {
	var gateways []Gateway
	Db.Order("id asc").Find(&gateways)
	return gateways
}

func GetGatewaysWithId(gatewayId string) []Gateway {
	var gateways []Gateway
	Db.Where("gateway_id = ?", gatewayId).Find(&gateways)
	return gateways
}

func GetOnlineGatewaysForNetwork(networkId string) []Gateway {

	if cacheGateways, ok := networkOnlineGatewaysCache.Get(networkId); ok {
		return cacheGateways.([]Gateway)
	}

	var gateways []Gateway
	Db.Where("network_id = ? AND last_heard > NOW() - INTERVAL '5 DAY'", networkId).Find(&gateways)

	// Store in cache
	networkOnlineGatewaysCache.Set(networkId, gateways, cache.DefaultExpiration)

	return gateways
}

func GetOnlineGatewaysForNetworkInBbox(networkId string, west float64, east float64, north float64, south float64) []Gateway {

	var gateways []Gateway
	Db.Where("network_id = ? AND last_heard > NOW() - INTERVAL '5 DAY'", networkId).
		Where("latitude >= ? AND latitude <= ? AND longitude >= ? AND longitude <= ?", south, north, west, east).
		Where("NOT (latitude = 0 AND longitude = 0)").
		Find(&gateways)

	return gateways
}

func GetGateway(indexer GatewayIndexer) (Gateway, error) {
	var gatewayDb Gateway
	i, ok := gatewayDbCache.Load(indexer)
	if ok {
		//log.Println("Gateway from cache")
		gatewayDb = i.(Gateway)
	} else {
		gatewayDb = Gateway{NetworkId: indexer.NetworkId, GatewayId: indexer.GatewayId}
		//log.Println("Gateway from DB")
		err := Db.First(&gatewayDb, &gatewayDb).Error
		if err != nil {
			return gatewayDb, err
		}
		if gatewayDb.ID != 0 {
			gatewayDbCache.Store(indexer, gatewayDb)
		}
	}
	return gatewayDb, nil
}

func GetGatewayLastMovedTime(networkId string, gatewayId string) time.Time {
	var movedTime time.Time
	lastMovedQuery := `
SELECT max(installed_at) FROM gateway_locations
WHERE network_id = ?
AND gateway_id = ?`
	timeRow := Db.Raw(lastMovedQuery, networkId, gatewayId).Row()
	timeRow.Scan(&movedTime)
	return movedTime
}

func GetAllOldNamingTtnV2Antennas() []Antenna {
	var antennas []Antenna
	Db.Where("network_id LIKE 'NS_TTN_V2://%' OR network_id LIKE 'NS_TTS_V3://ttnv2@000013'").Find(&antennas)
	return antennas
}

func FindAntenna(networkId string, gatewayId string, antennaIndex uint8) Antenna {
	var antenna Antenna
	antennaIndexer := AntennaIndexer{NetworkId: networkId, GatewayId: gatewayId, AntennaIndex: antennaIndex}

	i, ok := antennaDbCache.Load(antennaIndexer)
	if ok {
		antenna = i.(Antenna)
	} else {
		antenna = Antenna{NetworkId: networkId, GatewayId: gatewayId, AntennaIndex: antennaIndex}
		Db.FirstOrCreate(&antenna, &antenna)
	}
	return antenna
}

func GetAntennaForGateway(networkId string, gatewayId string) []Antenna {
	var antennas []Antenna
	Db.Where("network_id = ? and gateway_id = ?", networkId, gatewayId).Find(&antennas)
	return antennas
}

func GetAntennasForNetwork(networkId string) []Antenna {
	var antennas []Antenna
	Db.Where("network_id = ?", networkId).Find(&antennas)
	return antennas
}

func UpdatePacketsAntennaId(oldAntennaId uint, newAntennaId uint) {
	Db.Model(&Packet{}).Where("antenna_id = ?", oldAntennaId).Update("antenna_id", newAntennaId)
}

func GetDistinctGatewaysInLocations() []GatewayIndexer {
	var gateways []GatewayIndexer
	Db.Model(&GatewayLocation{}).Distinct("network_id", "gateway_id").Find(&gateways)
	return gateways
}

func GetGatewayLocations(networkId string, gatewayId string) []GatewayLocation {
	var locations []GatewayLocation
	Db.Where("network_id = ? and gateway_id = ?", networkId, gatewayId).Order("installed_at asc").Find(&locations)
	return locations
}

func DeleteGatewayLocations(networkId string, gatewayId string) {
	Db.Where("network_id = ? and gateway_id = ?", networkId, gatewayId).Delete(GatewayLocation{})
}

func InsertGatewayLocations(locations []GatewayLocation) {
	Db.Create(&locations)
}

func GetGridcellsForAntenna(antenna Antenna) []GridCell {
	var gridCells []GridCell
	Db.Where("antenna_id = ?", antenna.ID).Find(&gridCells)
	return gridCells
}

func GetGridCell(indexer GridCellIndexer) (GridCell, error) {
	var gridCell GridCell
	gridCell.AntennaID = indexer.AntennaID
	gridCell.X = indexer.X
	gridCell.Y = indexer.Y
	err := Db.FirstOrCreate(&gridCell, &gridCell).Error
	return gridCell, err
}

func SaveGridCell(gridCell GridCell) {
	Db.Save(&gridCell)
}

func CreateGridCells(gridCells []GridCell) error {
	// On conflict override
	tx := Db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&gridCells)
	return tx.Error
}

func DeleteGridCellsForAntenna(antenna Antenna) {
	Db.Where(&GridCell{AntennaID: antenna.ID}).Delete(&GridCell{})
}

func GetRadarBeam(indexer RadarBeamIndexer) (RadarBeam, error) {
	var radarBeam RadarBeam
	radarBeam.AntennaID = indexer.AntennaID
	radarBeam.Level = indexer.Level
	radarBeam.Bearing = indexer.Bearing
	err := Db.FirstOrCreate(&radarBeam, &radarBeam).Error
	return radarBeam, err
}

func SaveRadarBeam(radarBeam RadarBeam) {
	Db.Save(&radarBeam)
}

func GetRadarBeamsForAntenna(antenna Antenna) []RadarBeam {
	var radarBeams []RadarBeam
	Db.Where("antenna_id = ?", antenna.ID).Find(&radarBeams)
	return radarBeams
}

func CreateRadarBeams(radarBeams []RadarBeam) error {
	// On conflict override
	tx := Db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&radarBeams)
	return tx.Error
}

func DeleteRadarBeamsForAntenna(antenna Antenna) {
	Db.Where(&RadarBeam{AntennaID: antenna.ID}).Delete(&RadarBeam{})
}
