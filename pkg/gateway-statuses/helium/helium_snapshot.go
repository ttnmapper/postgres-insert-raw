package helium

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"ttnmapper-postgres-insert-raw/pkg/types"
)

func FetchSnapshot() ([]HotspotSnapshot, error) {
	// Fetch the latest snapshot which is stored at 01:00:00Z today
	today := time.Now()
	filename := fmt.Sprintf("%4d-%02d-%02d", today.Year(), today.Month(), today.Day())
	url := "https://snapshots.helium.wtf/mainnet/hotspots/network/" + filename + ".json.gz"
	log.Println("Fetching Helium Snapshot", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	var statuses []HotspotSnapshot
	err = json.NewDecoder(reader).Decode(&statuses)
	if err != nil {
		return nil, err
	}

	return statuses, nil
}

func HeliumHotspotSnapshotToTtnMapperGateway(hotspot HotspotSnapshot) (types.TtnMapperGateway, error) {

	gateway := types.TtnMapperGateway{
		NetworkId:    "NS_HELIUM://000024",
		GatewayId:    hotspot.Address,
		GatewayEui:   "",
		Name:         hotspot.Name,
		AntennaIndex: 0,
		//Time:         hotspot.Status.Timestamp.UnixNano(),
		//Timestamp:                   0,
		//FineTimestamp:               0,
		//FineTimestampEncrypted:      nil,
		//FineTimestampEncryptedKeyId: "",
		//ChannelIndex:                hotspot.Channel,
		//Rssi:                        hotspot.Rssi,
		//SignalRssi:                  0,
		//Snr:                         hotspot.Snr,
		Latitude:  hotspot.Latitude,
		Longitude: hotspot.Longitude,
		//Altitude:         hotspot.Elevation,
		LocationAccuracy: 0,
		LocationSource:   "",
	}

	// Only accept last heard times for online gateways
	if hotspot.Online == "online" {
		gateway.Time = time.Now().UnixNano()
	} else {
		gateway.Time = 0
	}

	gateway.Attributes = make(map[string]interface{}, 0)
	gateway.Attributes["mode"] = hotspot.Mode
	// Add other fields as required

	return gateway, nil
}
