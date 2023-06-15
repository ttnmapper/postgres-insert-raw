package helium

import "time"

type HotspotApiResponse struct {
	Data   []Hotspot `json:"data"`
	Cursor string    `json:"cursor"`
}

/*
	{
	   "lng":-78.74710984473072,
	   "lat":42.99533035889439,
	   "timestamp_added":"2021-11-13T17:29:11.000000Z",
	   "status":{
	      "timestamp":null,
	      "online":"online",
	      "listen_addrs":null,
	      "height":null
	   },
	   "reward_scale":null,
	   "payer":"13ENbEQPAvytjLnqavnbSAzurhGoCSNkGECMx7eHHDAfEaDirdY",
	   "owner":"135SUwbieAFXYXEbZbNZdHW8U9kAkyRCAguojRp3Sv2reWjsyhd",
	   "nonce":1,
	   "name":"bubbly-honeysuckle-dolphin",
	   "mode":"full",
	   "location_hex":"882aa6d823fffff",
	   "location":"8c2aa6d823927ff",
	   "last_poc_challenge":null,
	   "last_change_block":1097250,
	   "geocode":{
	      "short_street":"Presidents Walk",
	      "short_state":"NY",
	      "short_country":"US",
	      "short_city":"Buffalo",
	      "long_street":"Presidents Walk",
	      "long_state":"New York",
	      "long_country":"United States",
	      "long_city":"Buffalo",
	      "city_id":"YnVmZmFsb25ldyB5b3JrdW5pdGVkIHN0YXRlcw"
	   },
	   "gain":90,
	   "elevation":5,
	   "block_added":1097250,
	   "block":1097251,
	   "address":"11j4jiWSv44uJG351X7dXyXvrfPNoio68Jggo7Qu2RCeNNS4rET"
	}

	{
		"lng": 29.790906981600887,
		"lat": 40.70157817240036,
		"timestamp_added": "2022-01-06T10:30:16Z",
		"status": {
			"timestamp": "2022-01-06T19:52:56.186Z",
			"online": "online",
			"listen_addrs": [
				"/ip4/185.180.29.146/tcp/44158"
			],
			"height": 1089481
		},
		"reward_scale": 1,
		"payer": "134C7Hn3vhfBLQZex4PVwtxQ2uPJH97h9YD2bhzy1W2XhMJyY6d",
		"owner": "14gruvstK6d5BXqtggivFL7JD3N6DoWCK3PHZTThe1Kjptmq6Am",
		"nonce": 1,
		"name": "elegant-syrup-kitten",
		"mode": "full",
		"location_hex": "881ec924a9fffff",
		"location": "8c1ec924a8657ff",
		"last_poc_challenge": 1169857,
		"last_change_block": 1170211,
		"geocode": {
			"city_id": "a29jYWVsaXR1cmtleQ",
			"long_city": null,
			"long_country": "Turkey",
			"long_state": "Kocaeli",
			"long_street": "Örcün Yolu Caddesi",
			"short_city": null,
			"short_country": "TR",
			"short_state": "Kocaeli",
			"short_street": "Örcün Yolu Cd."
		},
		"gain": 40,
		"elevation": 20,
		"block_added": 1168392,
		"block": 1170217,
		"address": "112W2BfRWFZq6Tvc1QAP6eCRrbfQFzJWurw6ay2ToMY9gjqTBQ6j"
	}
*/
type Hotspot struct {
	Longitude        float64       `json:"lng"`
	Latitude         float64       `json:"lat"`
	TimestampAdded   time.Time     `json:"timestamp_added"`
	Status           HotspotStatus `json:"status"`
	RewardScale      interface{}   `json:"reward_scale"`
	Payer            string        `json:"payer"`
	Owner            string        `json:"owner"`
	Nonce            int           `json:"nonce"`
	Name             string        `json:"name"`
	Mode             string        `json:"mode"`
	LocationHex      string        `json:"location_hex"`
	Location         string        `json:"location"`
	LastPocChallenge interface{}   `json:"last_poc_challenge"`
	LastChangeBlock  int           `json:"last_change_block"`
	Geocode          interface{}   `json:"geocode"`
	Gain             int           `json:"gain"`
	Elevation        int32         `json:"elevation"`
	BlockAdded       int           `json:"block_added"`
	Block            int           `json:"block"`
	Address          string        `json:"address"`
}

/*
	"status":{
	   "timestamp":null,
	   "online":"online",
	   "listen_addrs":null,
	   "height":null
	},
*/
type HotspotStatus struct {
	Timestamp   time.Time   `json:"timestamp"`
	Online      string      `json:"online"`
	ListenAddrs interface{} `json:"listen_addrs"`
	Height      interface{} `json:"height"`
}

type HotspotSnapshot struct {
	Address      string  `json:"address"`
	Mode         string  `json:"mode"`
	Owner        string  `json:"owner"`
	Location     string  `json:"location"`
	Name         string  `json:"name"`
	Online       string  `json:"online"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"lng"`
	ShortStreet  string  `json:"short_street"`
	ShortCity    string  `json:"short_city"`
	ShortState   string  `json:"short_state"`
	ShortCountry string  `json:"short_country"`
}

type Disk91Snapshot struct {
	HotspotId  string `json:"hotspotId"`
	AnimalName string `json:"animalName"`
	Position   struct {
		LastDatePosition int64   `json:"lastDatePosition"`
		Lat              float64 `json:"lat"`
		Lng              float64 `json:"lng"`
		Country          string  `json:"country"`
		City             string  `json:"city"`
		Alt              float64 `json:"alt"`
		Gain             float64 `json:"gain"`
	} `json:"position"`
	LastSeen int64 `json:"lastSeen"`
}
