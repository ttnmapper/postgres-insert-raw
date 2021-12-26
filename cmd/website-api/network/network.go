package network

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{network_id}/gateways", GetGateways)
	router.Get("/{network_id}/gateways/{page}", GetGatewaysPaged)

	return router
}

func GetGateways(writer http.ResponseWriter, request *http.Request) {
	var err error

	networkId := chi.URLParam(request, "network_id")
	networkId, err = url.PathUnescape(networkId)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	dbGateways := database.GetOnlineGatewaysForNetwork(networkId)
	responseGateways := responses.DbGatewaysToResponse(dbGateways)
	render.JSON(writer, request, responseGateways)

	//for _, gateway := range responseGateways {
	//	render.JSON(writer, request, gateway)
	//	//writer.Write([]byte("\n"))
	//}
}

func GetGatewaysPaged(writer http.ResponseWriter, request *http.Request) {
	var err error
	var pageLimit = 10000

	networkId := chi.URLParam(request, "network_id")
	networkId, err = url.PathUnescape(networkId)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	page := chi.URLParam(request, "page")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err.Error())
		pageInt = 0
	}

	dbGateways := database.GetOnlineGatewaysForNetwork(networkId)
	pageStart := pageInt * pageLimit
	pageEnd := pageStart + pageLimit
	log.Println(len(dbGateways), pageStart, pageEnd)
	if pageStart > len(dbGateways) {
		pageStart = len(dbGateways)
	}
	if pageEnd > len(dbGateways) {
		pageEnd = len(dbGateways)
	}

	dbGateways = dbGateways[pageStart:pageEnd]
	responseGateways := responses.DbGatewaysToResponse(dbGateways)
	render.JSON(writer, request, responseGateways)
}
