package helium

import (
	"log"
	"testing"
	"ttnmapper-postgres-insert-raw/pkg/utils"
)

func TestFetchSnapshot(t *testing.T) {
	snapshotTime, hotspots, err := FetchSnapshot()
	if err != nil {
		t.Fatalf(err.Error())
	}

	log.Printf("Found %d hotspots in snapshot", len(hotspots))
	for _, hotspot := range hotspots {
		gateway, err := HotspotSnapshotToTtnMapperGateway(snapshotTime, hotspot)
		if err != nil {
			t.Fatalf(err.Error())
		}
		log.Println(utils.PrettyPrint(gateway))
		break
	}
}
