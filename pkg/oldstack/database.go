package oldstack

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var (
	db *gorm.DB
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		databaseContext.User, databaseContext.Password, databaseContext.Host, databaseContext.Port, databaseContext.Database)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	})
	if err != nil {
		panic(err.Error())
	}
}

func GetGatewayUpdates(limit int, offset int) []GatewayUpdate {
	var gatewayUpdates []GatewayUpdate
	db.Limit(limit).Offset(offset).Find(&gatewayUpdates)
	return gatewayUpdates
}

func GetPackets(limit int, offset uint) []Packet {
	var packets []Packet
	//db.Limit(limit).Offset(offset).Find(&packets)
	db.Limit(limit).Where("id > ?", offset).Find(&packets)
	return packets
}

func GetExperimentPackets(limit int, offset int) []Packet {
	var packets []Packet
	db.Table("experiments").Limit(limit).Offset(offset).Find(&packets)
	return packets
}

func GetPacket(id int) Packet {
	var packet Packet
	packet.ID = uint(id)
	db.First(&packet, &packet)
	return packet
}
