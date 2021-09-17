package oldstack

import (
	"log"
	"testing"
)

func TestGwaddrToNetIdEui(t *testing.T) {
	var (
		networkId  string
		gatewayId  string
		gatewayEui string
	)
	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("0011223344556677")
	if gatewayId != "eui-0011223344556677" {
		t.Fatalf("%s != \"eui-0011223344556677\"", gatewayId)
	}
	if gatewayEui != "0011223344556677" {
		t.Fatalf("%s != \"0011223344556677\"", gatewayEui)
	}

	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("eui-0011223344556677")
	if gatewayId != "eui-0011223344556677" {
		t.Fatalf("%s != \"eui-0011223344556677\"", gatewayId)
	}
	if gatewayEui != "0011223344556677" {
		t.Fatalf("%s != \"0011223344556677\"", gatewayEui)
	}

	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("aabbccddeeff1122")
	if gatewayId != "eui-aabbccddeeff1122" {
		t.Fatalf("%s != \"eui-aabbccddeeff1122\"", gatewayId)
	}
	if gatewayEui != "AABBCCDDEEFF1122" {
		t.Fatalf("%s != \"AABBCCDDEEFF1122\"", gatewayEui)
	}

	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("eui-aabbccddeeff1122")
	if gatewayId != "eui-aabbccddeeff1122" {
		t.Fatalf("%s != \"eui-aabbccddeeff1122\"", gatewayId)
	}
	if gatewayEui != "AABBCCDDEEFF1122" {
		t.Fatalf("%s != \"AABBCCDDEEFF1122\"", gatewayEui)
	}

	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("hello-my-gateway")
	if gatewayId != "hello-my-gateway" {
		t.Fatalf("%s != \"hello-my-gateway\"", gatewayId)
	}
	if gatewayEui != "" {
		t.Fatalf("%s != \"\"", gatewayEui)
	}
	if networkId != "thethingsnetwork.org" {
		t.Fatalf("%s != \"thethingsnetwork.org\"", networkId)
	}

	networkId, gatewayId, gatewayEui = GwaddrToNetIdEui("dragino-pg1301-b827eb752467ffff")
	log.Println(networkId, gatewayId, gatewayEui)
}

func TestDatarateToSfBw(t *testing.T) {
	sf, bw := DatarateToSfBw("SF7BW125")
	if sf != 7 {
		t.Fatalf("SF %d != 7", sf)
	}
	if bw != 125000 {
		t.Fatalf("%d != 125000", bw)
	}

	sf, bw = DatarateToSfBw("")
	if sf != 7 {
		t.Fatalf("SF %d != 7", sf)
	}
	if bw != 125000 {
		t.Fatalf("%d != 125000", bw)
	}

	sf, bw = DatarateToSfBw("SF7BW250")
	if sf != 7 {
		t.Fatalf("SF %d != 7", sf)
	}
	if bw != 250000 {
		t.Fatalf("%d != 250000", bw)
	}

	sf, bw = DatarateToSfBw("SF12BW125")
	if sf != 12 {
		t.Fatalf("SF %d != 12", sf)
	}
	if bw != 125000 {
		t.Fatalf("%d != 125000", bw)
	}
}
