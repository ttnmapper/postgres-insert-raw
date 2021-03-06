package types

import (
	"time"
)

type Packet struct {
	ID   uint
	Time time.Time `gorm:"not null;index:idx_packets_antenna_id_time,priority:2;index:idx_packets_device_id_time,priority:2"` // index priority 11 is lower than default 10. Device and gateway is less unique, so will filter better first step.

	DeviceID uint `gorm:"not null;index:idx_packets_device_id_time,priority:1"`

	FPort uint8
	FCnt  uint32

	FrequencyID  uint
	DataRateID   uint
	CodingRateID uint

	// Gateway data
	AntennaID              uint `gorm:"not null;index:idx_packets_antenna_id_time,priority:1;index:idx_packets_antenna_id_latitude,priority:1;index:idx_packets_antenna_id_longitude,priority:1"`
	GatewayTime            *time.Time
	Timestamp              *uint32
	FineTimestamp          *uint64
	FineTimestampEncrypted *[]byte
	FineTimestampKeyID     *uint
	ChannelIndex           uint32
	Rssi                   float32  `gorm:"type:numeric(6,2)"`
	SignalRssi             *float32 `gorm:"type:numeric(6,2)"`
	Snr                    float32  `gorm:"type:numeric(5,2)"`

	Latitude         float64  `gorm:"not null;type:numeric(10,6);index:idx_packets_antenna_id_latitude,priority:2"`
	Longitude        float64  `gorm:"not null;type:numeric(10,6);index:idx_packets_antenna_id_longitude,priority:2"`
	Altitude         float64  `gorm:"type:numeric(6,1)"`
	AccuracyMeters   *float64 `gorm:"type:numeric(6,2)"`
	Satellites       *int32
	Hdop             *float64 `gorm:"type:numeric(4,1)"`
	AccuracySourceID uint

	ExperimentID *uint

	UserID      uint
	UserAgentID uint

	DeletedAt *time.Time
}

type Device struct {
	ID      uint
	AppId   string `gorm:"index:app_device_eui,unique"`
	DevId   string `gorm:"index:app_device_eui,unique"`
	DevEui  string `gorm:"index:app_device_eui,unique"`
	Packets []Packet
}

type Frequency struct {
	ID      uint
	Herz    uint64 `gorm:"unique;not null"`
	Packets []Packet
}

type DataRate struct {
	ID              uint
	Modulation      string `gorm:"index:data_rate,unique"` // LORA or FSK or LORA-E
	Bandwidth       uint64 `gorm:"index:data_rate,unique"`
	SpreadingFactor uint8  `gorm:"index:data_rate,unique"`
	Bitrate         uint64 `gorm:"index:data_rate,unique"`
	Packets         []Packet
}

type CodingRate struct {
	ID      uint
	Name    string `gorm:"unique;not null"`
	Packets []Packet
}

type AccuracySource struct {
	ID      uint
	Name    string `gorm:"unique;not null"`
	Packets []Packet
}

type Experiment struct {
	ID      uint
	Name    string `gorm:"unique;not null"`
	Packets []Packet
}

type User struct {
	ID         uint
	Identifier string `gorm:"unique;not null"`
	Packets    []Packet
}

type UserAgent struct {
	ID      uint
	Name    string `gorm:"unique;not null"`
	Packets []Packet
}

// TODO: Currently we identify a gateway using the gateway ID provided by the network.
// But how are we going to identify them between networks, when data is sent via the packet broker?

type Antenna struct {
	ID uint

	// TTN gateway ID along with the Antenna index identifies a unique coverage area.
	NetworkId    string `gorm:"type:text;index:idx_gtw_id_antenna,unique"`
	GatewayId    string `gorm:"type:text;index:idx_gtw_id_antenna,unique"`
	AntennaIndex uint8  `gorm:"index:idx_gtw_id_antenna,unique"`

	// For now we do not set antenna locations, but add it here for future use
	//Latitude         *float64
	//Longitude        *float64
	//Altitude         *int32

	Packets []Packet
}

type Gateway struct {
	ID uint

	NetworkId   string `gorm:"type:text;UNIQUE_INDEX:idx_gtw_id"`
	GatewayId   string `gorm:"type:text;UNIQUE_INDEX:idx_gtw_id"`
	GatewayEui  *string
	Description *string

	Latitude         *float64
	Longitude        *float64
	Altitude         *int32
	LocationAccuracy *int32
	LocationSource   *string

	//AtLocationSince	time.Time // This value gets updated when the gateway moves
	LastHeard time.Time // This value always gets updated to reflect that the gateway is working

	//Antennas         []Antenna
	//GatewayLocations []GatewayLocation
}

type GatewayLocation struct {
	ID        uint
	NetworkId string `gorm:"type:text;INDEX:idx_gtw_id_install"`
	GatewayId string `gorm:"type:text;INDEX:idx_gtw_id_install"`

	InstalledAt time.Time `gorm:"INDEX:idx_gtw_id_install"`
	Latitude    float64
	Longitude   float64
}

// To blacklist a gateway set its location to 0,0
type GatewayLocationForce struct {
	ID        uint
	NetworkId string `gorm:"type:text;UNIQUE_INDEX:idx_gtw_id"`
	GatewayId string `gorm:"type:text;UNIQUE_INDEX:idx_gtw_id"`

	Latitude  float64
	Longitude float64
}

type FineTimestampKeyID struct {
	ID                          uint
	FineTimestampEncryptedKeyId string
}

// Indexers: These structs are the same as the ones above, but used to index the cache maps
type DeviceIndexer struct {
	DevId  string
	AppId  string
	DevEui string
}

type GatewayIndexer struct {
	NetworkId string
	GatewayId string
}

type AntennaIndexer struct {
	NetworkId    string
	GatewayId    string
	AntennaIndex uint8
}

type DataRateIndexer struct {
	Modulation      string // LORA or FSK or LORA-E
	Bandwidth       uint64
	SpreadingFactor uint8
	Bitrate         uint64
}
