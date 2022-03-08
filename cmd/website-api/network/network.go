package network

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/j4/gosm"
	"github.com/patrickmn/go-cache"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

var (
	// Caches for dynamic content, needs to expire
	networkZ5GatewaysCache *cache.Cache
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{network_id}/gateways", GetGateways)
	router.Get("/{network_id}/gateways/{page}", GetGatewaysPaged)
	router.Get("/{network_id}/gateways/z5tile/{x}/{y}", GetGatewaysInZ5Tile)

	networkZ5GatewaysCache = cache.New(4*time.Hour, 1*time.Hour)

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

	// Include all gateways from networks that feed data into this network via the packet broker
	peeredNetworks := database.GetPeeredNetworks(networkId)

	responseGateways := make([]responses.Gateway, 0)

	selfInPeeredNetworks := false
	for _, network := range peeredNetworks {
		if network.ForwarderNetworkId == networkId {
			selfInPeeredNetworks = true
		}
		dbGateways := database.GetOnlineGatewaysForNetwork(network.ForwarderNetworkId)
		responseGateways = append(responseGateways, responses.DbGatewaysToResponse(dbGateways)...)
	}

	// If not routing policies exist for this network, no gateways will be returned. So fetch for this specific network then.
	if !selfInPeeredNetworks {
		dbGateways := database.GetOnlineGatewaysForNetwork(networkId)
		responseGateways = append(responseGateways, responses.DbGatewaysToResponse(dbGateways)...)
	}

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

func GetGatewaysInZ5Tile(writer http.ResponseWriter, request *http.Request) {
	var err error

	networkId := chi.URLParam(request, "network_id")
	networkId, err = url.PathUnescape(networkId)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	x, err := strconv.Atoi(chi.URLParam(request, "x"))
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}
	y, err := strconv.Atoi(chi.URLParam(request, "y"))
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	cacheIndex := networkId + "/" + strconv.Itoa(x) + "/" + strconv.Itoa(y)
	if cacheGateways, ok := networkZ5GatewaysCache.Get(cacheIndex); ok {
		render.JSON(writer, request, cacheGateways.([]responses.Gateway))
		return
	}

	tile := gosm.NewTileWithXY(x, y, 5)
	north := tile.Lat
	west := tile.Long
	tile = gosm.NewTileWithXY(x+1, y+1, 5)
	south := tile.Lat
	east := tile.Long

	dbGateways := database.GetOnlineGatewaysForNetworkInBbox(networkId, west, east, north, south)
	responseGateways := responses.DbGatewaysToResponse(dbGateways)

	networkZ5GatewaysCache.Set(cacheIndex, responseGateways, cache.DefaultExpiration)

	render.JSON(writer, request, responseGateways)
}
