package thethingsstack

import "time"

type Gateway struct {
	Ids struct {
		GatewayId string `json:"gateway_id"`
		Eui       string `json:"eui"`
	} `json:"ids"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   time.Time         `json:"deleted_at"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Attributes  map[string]string `json:"attributes"`
	ContactInfo []interface{}     `json:"contact_info"`
	VersionIds  struct {
	} `json:"version_ids"`
	GatewayServerAddress string   `json:"gateway_server_address"`
	AutoUpdate           bool     `json:"auto_update"`
	UpdateChannel        string   `json:"update_channel"`
	FrequencyPlanId      string   `json:"frequency_plan_id"`
	FrequencyPlanIds     []string `json:"frequency_plan_ids"`
	Antennas             []struct {
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
			Altitude  int32   `json:"altitude"`
			Accuracy  int32   `json:"accuracy"`
			Source    string  `json:"source"`
		} `json:"location"`
		Placement string `json:"placement"`
	} `json:"antennas,omitempty"`
	StatusPublic             bool   `json:"status_public"`
	LocationPublic           bool   `json:"location_public"`
	ScheduleDownlinkLate     bool   `json:"schedule_downlink_late"`
	EnforceDutyCycle         bool   `json:"enforce_duty_cycle"`
	DownlinkPathConstraint   string `json:"downlink_path_constraint"`
	ScheduleAnytimeDelay     string `json:"schedule_anytime_delay"`
	UpdateLocationFromStatus bool   `json:"update_location_from_status"`
	LbsLnsSecret             struct {
	} `json:"lbs_lns_secret"`
	ClaimAuthenticationCode struct {
	} `json:"claim_authentication_code"`
	TargetCupsUri string `json:"target_cups_uri"`
	TargetCupsKey struct {
	} `json:"target_cups_key"`
	RequireAuthenticatedConnection bool `json:"require_authenticated_connection"`
	Lrfhss                         struct {
	} `json:"lrfhss"`
	DisablePacketBrokerForwarding bool `json:"disable_packet_broker_forwarding"`
}

// curl --header "Authorization: Bearer NNSXS.xxx" https://jpmeijers.eu1.cloud.thethings.industries/api/v3/gateways?field_mask=name,description,antennas
// curl --header "Authorization: Bearer NNSXS.xxx" https://jpmeijers.eu1.cloud.thethings.industries/api/v3/gateways?field_mask=ids,created_at,updated_at,deleted_at,name,description,attributes,contact_info,version_ids,gateway_server_address,auto_update,frequency_plan_id,frequency_plan_ids,antennas,status_public,location_public,schedule_downlink_late,enforce_duty_cycle,downlink_path_constraint,schedule_anytime_delay,update_location_from_status,lbs_lns_secret,claim_authentication_code,target_cups_uri,target_cups_key,require_authenticated_connection,lrfhss,disable_packet_broker_forwarding
type V3Gateways struct {
	Gateways []Gateway `json:"gateways"`
}

// curl --header "Authorization: Bearer NNSXS.xxx" https://packetworx.au1.cloud.thethings.industries/api/v3/gs/gateways/johann-non-basic-station-test/connection/stats
type Status struct {
	ConnectedAt          time.Time `json:"connected_at"`
	Protocol             string    `json:"protocol"`
	LastStatusReceivedAt time.Time `json:"last_status_received_at"`
	LastStatus           struct {
		Time     time.Time `json:"time"`
		Versions struct {
			TtnLwGatewayServer string `json:"ttn-lw-gateway-server"`
		} `json:"versions"`
		Ip      []string `json:"ip"`
		Metrics struct {
			Rxfw int `json:"rxfw"`
			Ackr int `json:"ackr"`
			Txin int `json:"txin"`
			Txok int `json:"txok"`
			Rxin int `json:"rxin"`
			Rxok int `json:"rxok"`
		} `json:"metrics"`
	} `json:"last_status"`
	LastUplinkReceivedAt   time.Time `json:"last_uplink_received_at"`
	UplinkCount            string    `json:"uplink_count"`
	LastDownlinkReceivedAt time.Time `json:"last_downlink_received_at"`
	DownlinkCount          string    `json:"downlink_count"`
	RoundTripTimes         struct {
		Min    string `json:"min"`
		Max    string `json:"max"`
		Median string `json:"median"`
		Count  int    `json:"count"`
	} `json:"round_trip_times"`
	SubBands []struct {
		MinFrequency             string  `json:"min_frequency"`
		MaxFrequency             string  `json:"max_frequency"`
		DownlinkUtilizationLimit float64 `json:"downlink_utilization_limit"`
		DownlinkUtilization      float64 `json:"downlink_utilization,omitempty"`
	} `json:"sub_bands"`
}
