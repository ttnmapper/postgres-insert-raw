package gateway

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"net/url"
	"time"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{network_id}/{gateway_id}/details", GetGatewayDetails)
	router.Get("/data", GetGatewayData)
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

func GetGatewayData(writer http.ResponseWriter, request *http.Request) {
	var err error

	gatewayId := request.URL.Query().Get("gateway_id")
	if gatewayId == "" {
		responses.RenderError(writer, request, errors.New("gateway_id not set"))
		return
	}
	networkId := request.URL.Query().Get("network_id")

	startTimeString := request.URL.Query().Get("start_time")
	startTime := time.Time{}
	if startTimeString != "" {
		// parse rfc-3339 datetime
		startTime, err = time.Parse(time.RFC3339, startTimeString)
		if err != nil {
			responses.RenderError(writer, request, errors.New("can not parse start_time"))
			return
		}
	}

	endTimeString := request.URL.Query().Get("end_time")
	endTime := time.Now()
	if endTimeString != "" {
		// parse rfc-3339 datetime
		endTime, err = time.Parse(time.RFC3339, endTimeString)
		if err != nil {
			responses.RenderError(writer, request, errors.New("can not parse end_time"))
			return
		}
	}

	gatewayDataDbRows, err := database.GetPacketsForGateway(networkId, gatewayId, startTime, endTime, 10000)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	result := make([]responses.DeviceMeasurement, 0)
	for gatewayDataDbRows.Next() {
		measurement := responses.DeviceMeasurement{}
		err = database.Db.ScanRows(gatewayDataDbRows, &measurement)
		if err != nil {
			responses.RenderError(writer, request, err)
			break
		}
		result = append(result, measurement)
	}

	render.JSON(writer, request, result)
}
