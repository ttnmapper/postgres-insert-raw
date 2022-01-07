package thethingsstack

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

func TestFetchStatuses(t *testing.T) {
	//tenantId := "jpmeijers"
	//apiKey := "NNSXS.7KYTHCV27ZWVIUHJCEA2X3GE3JMRRB3UFEZXXWY.CJRCATPOYFGSHKSSBJLKDIVX2HDHZKDJHYZSUGWY75TOL4CYMQOQ"
	tenantId := "packetworx"
	apiKey := "NNSXS.MVERCW4V6PM4VAELEQNHX3YX5MZSEINNKWO7NEY.NBWE6ENQZLOISGCFHYCPTZLDZTU5V6EH4LHXGFBJHWBKS67CIXAQ"
	gateways, err := FetchGateways(tenantId, apiKey)
	if err != nil {
		log.Println(err.Error())
	}
	//log.Println(utils.PrettyPrint(gateways))

	for _, gateway := range gateways {
		log.Println(utils.PrettyPrint(gateway.Antennas))
		continue

		status, err := FetchStatus(gateway, apiKey)
		if err != nil {
			log.Println(err.Error())
		}
		//log.Println(gateway.Ids.GatewayId, lastHeard)

		ttnMapperGateway, err := TtsApiGatewayToTtnMapperGateway(tenantId, gateway, status)
		if err != nil {
			log.Println(err)
		}
		//log.Println(utils.PrettyPrint(ttnMapperGateway))
		log.Println(ttnMapperGateway.GatewayId, ttnMapperGateway.Latitude, ttnMapperGateway.Longitude)
	}
}

func TestParseV3GatewayList(t *testing.T) {
	buf, err := ioutil.ReadFile("gateway-list-example.json")
	if err != nil {
		t.Fatalf(err.Error())
	}

	var apiResponse V3Gateways
	err = json.Unmarshal(buf, &apiResponse)
	if err != nil {
		t.Fatalf(err.Error())
	}

	//for i, ttsGateway := range apiResponse.Gateways {
	//	log.Println(i)
	//	//log.Println(utils.PrettyPrint(ttsGateway))
	//	gateway, _ := TtsApiGatewayToTtnMapperGateway(ttsGateway)
	//	log.Println(utils.PrettyPrint(gateway))
	//}
}
