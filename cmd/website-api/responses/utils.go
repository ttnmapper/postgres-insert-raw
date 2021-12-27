package responses

import (
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
		GatewayId:  gateway.GatewayId,
		NetworkId:  gateway.NetworkId,
		LastHeard:  gateway.LastHeard,
		Latitude:   gateway.Latitude,
		Longitude:  gateway.Longitude,
		Altitude:   gateway.Altitude,
	}

	if gateway.Description != nil {
		responseGw.Description = *gateway.Description
	}
	if gateway.GatewayEui != nil {
		responseGw.GatewayEUI = *gateway.GatewayEui
	}

	return responseGw
}
