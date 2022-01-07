package web

import "time"

// This is not final. This version obtained from @htdvisser
type WebGateway struct {
	ID          string   `json:"id"`
	Network     string   `json:"network"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Owner       string   `json:"owner,omitempty"`
	Owners      []string `json:"owners,omitempty"`
	Location    struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Altitude  int64   `json:"altitude"`
	} `json:"location,omitempty"`
	CountryCode string            `json:"country_code,omitempty"`
	Attributes  map[string]string `json:"attributes,omitempty"`
	Online      bool              `json:"online"`
	LastSeen    *time.Time        `json:"last_seen,omitempty"`
}
