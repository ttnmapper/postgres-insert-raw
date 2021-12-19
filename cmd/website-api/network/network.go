package network

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/gateways/{network_id}", GetGateways)

	return router
}

func GetGateways(writer http.ResponseWriter, request *http.Request) {
	errorResponse := responses.ErrorResponse{}
	var err error

	networkId := chi.URLParam(request, "network_id")
	networkId, err = url.PathUnescape(networkId)
	if err != nil {
		errorResponse.Success = false
		errorResponse.Message = err.Error()
		render.JSON(writer, request, errorResponse)
		return
	}

	dbGateways := database.GetOnlineGatewaysForNetwork(networkId)
	responseGateways := DbGatewaysToResponse(dbGateways)
	render.JSON(writer, request, responseGateways)
}

func DbGatewaysToResponse(dbGateways []database.Gateway) []responses.Gateway {
	var responseGateways []responses.Gateway

	for _, gateway := range dbGateways {
		responseGw := responses.Gateway{
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
		responseGateways = append(responseGateways, responseGw)
	}

	return responseGateways
}
