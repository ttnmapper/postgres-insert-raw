package utils

import (
	"errors"
	"math"
	"ttnmapper-postgres-insert-raw/pkg/types"
)

func CheckData(packet types.TtnMapperUplinkMessage) error {

	// Accuracy too low when satellites less than 4
	if !IsZeroOfUnderlyingType(packet.Satellites) {
		if packet.Satellites < 4 {
			return errors.New("less than 4 satellites")
		}
	}

	// Accuracy value must be below 10 metre
	if !IsZeroOfUnderlyingType(packet.AccuracyMeters) {
		if packet.AccuracyMeters > 10 {
			return errors.New("accuracy too low")
		}
	}

	// HDOP must be lower than 5
	if !IsZeroOfUnderlyingType(packet.Hdop) {
		if packet.Hdop > 5 {
			return errors.New("hdop is too high")
		}
	}

	// Latitude
	if IsZeroOfUnderlyingType(packet.Latitude) {
		return errors.New("latitude not set")
	}
	if packet.Latitude >= 90 || packet.Latitude <= -90 {
		return errors.New("latitude out of range")
	}

	// Longitude
	if IsZeroOfUnderlyingType(packet.Longitude) {
		return errors.New("longitude not set")
	}
	if packet.Longitude >= 180 || packet.Longitude <= -180 {
		return errors.New("longitude out of range")
	}

	// Null island
	if packet.Longitude < 1 && packet.Longitude > -1 && packet.Latitude < 1 && packet.Latitude > -1 {
		return errors.New("not accepting coordinates on null island")
	}

	return nil
}

func SanitizeData(packet *types.TtnMapperUplinkMessage) {
	// clamp altitude to ground if not set
	if IsZeroOfUnderlyingType(packet.Altitude) {
		packet.Altitude = 0
	}

	// handle overflow for a small range just below 2^16
	if packet.Altitude > (math.Pow(2, 16)-1000) && packet.Altitude < math.Pow(2, 16) {
		packet.Altitude = packet.Altitude - math.Pow(2, 16) // Negative altitude
	}

	// Null island
	if packet.Longitude < 1 && packet.Longitude > -1 && packet.Latitude < 1 && packet.Latitude > -1 {
		packet.Latitude = 0
		packet.Longitude = 0
	}

	// Latitude
	if packet.Latitude >= 90 || packet.Latitude <= -90 {
		packet.Latitude = 0
	}

	// Longitude
	if packet.Longitude >= 180 || packet.Longitude <= -180 {
		packet.Longitude = 0
	}
}

// Some single channel gateways send frequency in Hz, not MHz
// TTNv3 also sends the frequency in Herz, not MHz like V2 - change to Hz here
// Below 1MHz assume the value is passed in MHz not Hz, so convert to Hz
func SanitizeFrequency(frequency float64) uint64 {
	// Before we stored the frequency (and pade.nl)
	if frequency == 9.999 {
		return 0
	}

	// 868.1 to 868100000 - but we will lose the decimals
	if frequency < 1000.0 {
		frequency = frequency * 1000000
	}

	// 868400000000000 to 868400000
	if frequency > 1000000000 {
		frequency = frequency / 1000000
	}

	// 869099976 to 869100000
	frequency = math.Round(frequency/1000) * 1000
	frequencyInt := uint64(frequency)

	return frequencyInt
}

func ValidateChirpNetworkAddress(address string) (err error) {
	if address == "" {
		return errors.New("network address is empty")
	}

	return
}
