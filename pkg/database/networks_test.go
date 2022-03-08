package database

import (
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

func TestGetOnlineGatewaysForNetwork(t *testing.T) {
	initDb()

	gateways := GetOnlineGatewaysForNetwork("NS_TTS_V3://packetworx@000013")
	log.Println(utils.PrettyPrint(gateways))
}
