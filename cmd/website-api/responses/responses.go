package responses

import "time"

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

//type NetworkGateways struct {
//	_ []Gateway
//}

type Gateway struct {
	DatabaseId  uint      `json:"database_id"`
	GatewayId   string    `json:"gateway_id"`
	Description string    `json:"description"`
	GatewayEUI  string    `json:"gateway_eui"`
	NetworkId   string    `json:"network_id"`
	LastHeard   time.Time `json:"last_heard"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Altitude    int32     `json:"altitude"`
}
