package gateway

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{network_id}/{gateway_id}/details", GetGatewayDetails)
	router.Get("/{network_id}/{gateway_id}/radar", GetGatewayRadar)

	return router
}

func GetGatewayDetails(writer http.ResponseWriter, request *http.Request) {
	var err error

	networkId := chi.URLParam(request, "network_id")
	networkId, err = url.PathUnescape(networkId)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}
	if networkId == "" {
		responses.RenderError(writer, request, errors.New("network_id not set"))
		return
	}

	gatewayId := chi.URLParam(request, "gateway_id")
	gatewayId, err = url.PathUnescape(gatewayId)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}
	if gatewayId == "" {
		responses.RenderError(writer, request, errors.New("gateway_id not set"))
		return
	}

	indexer := database.GatewayIndexer{NetworkId: networkId, GatewayId: gatewayId}
	gateway, err := database.GetGateway(indexer)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	responseGateway := responses.DbGatewayToResponse(gateway)
	render.JSON(writer, request, responseGateway)
}
