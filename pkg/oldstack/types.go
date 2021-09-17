package oldstack

import (
	"database/sql"
	"time"
)

type GatewayUpdate struct {
	ID             uint      `gorm:"column:id"`
	GatewayAddress string    `gorm:"column:gwaddr"`
	InstalledAt    time.Time `gorm:"column:datetime"`
	Latitude       float64   `gorm:"column:lat"`
	Longitude      float64   `gorm:"column:lon"`
	Altitude       float64   `gorm:"column:alt"`
	LastUpdate     time.Time `gorm:"column:last_update"`
}

type Packet struct {
	ID             uint            `gorm:"column:id"`
	Time           time.Time       `gorm:"column:time"`
	NodeAddr       string          `gorm:"column:nodeaddr"`
	AppId          sql.NullString  `gorm:"column:appeui"` // null
	GatewayAddr    string          `gorm:"column:gwaddr"`
	Modulation     sql.NullString  `gorm:"column:modulation"` // null
	DataRate       sql.NullString  `gorm:"column:datarate"`   // null
	Snr            float64         `gorm:"column:snr"`
	Rssi           float64         `gorm:"column:rssi"`
	Frequency      float64         `gorm:"column:freq"`   // seen 9.999
	FrameCount     sql.NullInt32   `gorm:"column:fcount"` // null
	Latitude       float64         `gorm:"column:lat"`
	Longitude      float64         `gorm:"column:lon"`
	Altitude       sql.NullFloat64 `gorm:"column:alt"`        // null
	Accuracy       sql.NullFloat64 `gorm:"column:accuracy"`   // null
	Hdop           sql.NullFloat64 `gorm:"column:hdop"`       // null
	Satellites     sql.NullInt32   `gorm:"column:sats"`       // null
	Provider       sql.NullString  `gorm:"column:provider"`   // null
	MqttTopic      sql.NullString  `gorm:"column:mqtt_topic"` // null
	UserAgent      sql.NullString  `gorm:"column:user_agent"` // null
	UserId         string          `gorm:"column:user_id"`
	ExperimentName string          `gorm:"column:name"`
}
