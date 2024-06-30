package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

func InsertPacket(packet *Packet) error {
	err := Db.Create(&packet).Error
	return err
}

func GetPacketsForAntennaAfter(antenna Antenna, afterTime time.Time) (*sql.Rows, error) {
	// Get all existing packets since gateway last moved
	return Db.Model(&Packet{}).Where("antenna_id = ? AND time > ? AND experiment_id IS NULL", antenna.ID, afterTime).Rows() // server side cursor
}

func GetPacketsForDevice(networkId string, applicationId string, deviceId string, startTime time.Time, endTime time.Time, limit *int64) (*sql.Rows, error) {
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

	session = session.Where("experiment_id IS NULL")
	session = session.Where("time > ? AND time < ?", startTime, endTime)
	if networkId != "" {
		session = session.Where("d.network_id = ?", networkId)
	}
	if applicationId != "" {
		session = session.Where("d.app_id = ?", applicationId)
	}
	if deviceId != "" {
		session = session.Where("d.dev_id = ?", deviceId)
	}

	if limit != nil {
		session = session.Limit(int(*limit))
	}

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

	session = session.Where("experiment_id IS NULL")
	session = session.Where("a.gateway_id = ?", gatewayId)
	session = session.Where("time > ? AND time < ?", startTime, endTime)
	if networkId != "" {
		session = session.Where("a.network_id = ?", networkId)
	}

	session = session.Limit(limit)

	return session.Rows()
}

func GetPacketsForExperiment(experiment string, startTime time.Time, endTime time.Time, limit int) (*sql.Rows, error) {
	session := Db.Model(&Packet{})
	session = session.Select("packets.id, packets.time, packets.f_port, packets.f_cnt, packets.gateway_time, fine_timestamp, packets.channel_index, packets.rssi, packets.signal_rssi, packets.snr, packets.latitude, packets.longitude, packets.altitude, packets.accuracy_meters, packets.satellites, packets.hdop, app_id, dev_id, dev_eui, d.network_id as device_network_id, f.herz as frequency, modulation, bandwidth, spreading_factor, bitrate, cr.name as coding_rate, a.network_id as gateway_network_id, gateway_id, antenna_index, \"as\".name as accuracy_source, ua.name as user_agent, e.name as experiment")
	session = session.Joins("JOIN experiments e on packets.experiment_id = e.id")
	session = session.Joins("JOIN devices d on packets.device_id = d.id")
	session = session.Joins("JOIN frequencies f on packets.frequency_id = f.id")
	session = session.Joins("JOIN data_rates dr on packets.data_rate_id = dr.id")
	session = session.Joins("JOIN coding_rates cr on packets.coding_rate_id = cr.id")
	session = session.Joins("JOIN antennas a on packets.antenna_id = a.id")
	session = session.Joins("JOIN accuracy_sources \"as\" on packets.accuracy_source_id = \"as\".id")
	session = session.Joins("JOIN user_agents ua on packets.user_agent_id = ua.id")
	session = session.Joins("JOIN users on packets.user_id = users.id")

	//session = session.Where("experiment_id IS NULL")
	session = session.Where("e.name = ?", experiment)
	session = session.Where("time > ? AND time < ?", startTime, endTime)

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
