package gateway

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"net/url"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/layers/radar"
)

func GetGatewayRadarMulti(w http.ResponseWriter, r *http.Request) {
	errorResponse := responses.ErrorResponse{}
	var err error

	networkId := chi.URLParam(r, "network_id")
	networkId, err = url.PathUnescape(networkId)
	gatewayId := chi.URLParam(r, "gateway_id")
	gatewayId, err = url.PathUnescape(gatewayId)
	log.Println(networkId, gatewayId)

	geoJson := radar.GenerateRadarMulti(networkId, gatewayId)
	if geoJson == nil {
		errorResponse.Success = false
		errorResponse.Message = "could not generate geojson"
		render.JSON(w, r, errorResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(geoJson)
	if err != nil {
		log.Println("could not write geojson to response")
		return
	}
}

func GetGatewayRadarSingle(w http.ResponseWriter, r *http.Request) {
	errorResponse := responses.ErrorResponse{}
	var err error

	networkId := chi.URLParam(r, "network_id")
	networkId, err = url.PathUnescape(networkId)
	gatewayId := chi.URLParam(r, "gateway_id")
	gatewayId, err = url.PathUnescape(gatewayId)
	log.Println(networkId, gatewayId)

	geoJson := radar.GenerateRadarSingle(networkId, gatewayId)
	if geoJson == nil {
		errorResponse.Success = false
		errorResponse.Message = "could not generate geojson"
		render.JSON(w, r, errorResponse)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(geoJson)
	if err != nil {
		log.Println("could not write geojson to response")
		return
	}
}
