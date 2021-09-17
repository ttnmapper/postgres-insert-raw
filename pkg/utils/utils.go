package utils

import (
	"encoding/json"
	"log"
	"math"
	"reflect"
)

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func IsZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func CapFloatTo(input float64, digits uint, decimals uint) float64 {
	maxValue := math.Pow(10, float64(digits)) - 1         // 9999
	maxValue = maxValue / math.Pow(10, float64(decimals)) // 99.99
	input = math.Min(input, maxValue)
	input = math.Max(input, 0-maxValue)
	return input
}
