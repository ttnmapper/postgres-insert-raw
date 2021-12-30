package device

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"time"
	"ttnmapper-postgres-insert-raw/cmd/website-api/responses"
	"ttnmapper-postgres-insert-raw/pkg/database"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/data", GetDeviceData)

	return router
}

func GetDeviceData(writer http.ResponseWriter, request *http.Request) {
	var err error

	deviceId := request.URL.Query().Get("dev_id")
	if deviceId == "" {
		responses.RenderError(writer, request, errors.New("dev_id not set"))
		return
	}

	applicationId := request.URL.Query().Get("app_id")
	networkId := request.URL.Query().Get("network_id")
	log.Println("network_id", networkId)

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

	deviceDataDbRows, err := database.GetPacketsForDevice(networkId, applicationId, deviceId, startTime, endTime, 10000)
	if err != nil {
		responses.RenderError(writer, request, err)
		return
	}

	result := make([]responses.DeviceMeasurement, 0)
	for deviceDataDbRows.Next() {
		measurement := responses.DeviceMeasurement{}
		err = database.Db.ScanRows(deviceDataDbRows, &measurement)
		if err != nil {
			responses.RenderError(writer, request, err)
			break
		}
		result = append(result, measurement)
	}

	render.JSON(writer, request, result)
}
