package responses

import (
	"encoding/json"
	"github.com/go-chi/render"
	"net/http"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func RenderError(writer http.ResponseWriter, request *http.Request, err error) {
	errorResponse := ErrorResponse{}
	errorResponse.Success = false
	errorResponse.Message = err.Error()
	render.JSON(writer, request, errorResponse)
}

func DbGatewaysToResponse(dbGateways []database.Gateway) []Gateway {
	responseGateways := make([]Gateway, 0)

	for _, gateway := range dbGateways {
		responseGw := DbGatewayToResponse(gateway)
		responseGateways = append(responseGateways, responseGw)
	}

	return responseGateways
}

func DbGatewayToResponse(gateway database.Gateway) Gateway {

	responseGw := Gateway{
		DatabaseId: gateway.ID,
		NetworkId:  gateway.NetworkId,
		GatewayId:  gateway.GatewayId,
		LastHeard:  gateway.LastHeard,
		Latitude:   gateway.Latitude,
		Longitude:  gateway.Longitude,
		Altitude:   gateway.Altitude,
	}

	if gateway.GatewayEui != nil {
		responseGw.GatewayEUI = *gateway.GatewayEui
	}
	if gateway.Name != nil {
		responseGw.Name = *gateway.Name
	}

	var attributes interface{}
	err := json.Unmarshal(gateway.Attributes, attributes)
	if err != nil {
		// nothing
	}
	responseGw.Attributes = attributes.(map[string]interface{})

	return responseGw
}
