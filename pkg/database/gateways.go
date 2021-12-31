package database

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

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
