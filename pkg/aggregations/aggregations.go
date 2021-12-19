package aggregations

import (
	"time"
	"ttnmapper-postgres-insert-raw/pkg/aggregations/grid_cell"
	"ttnmapper-postgres-insert-raw/pkg/aggregations/radar_beam"
	"ttnmapper-postgres-insert-raw/pkg/database"
	"ttnmapper-postgres-insert-raw/pkg/types"
)

func AggregateNewData(message types.TtnMapperUplinkMessage) {
	grid_cell.AggregateNewData(message)
	radar_beam.AggregateNewData(message)
}

func AggregateMovedGateway(message types.TtnMapperGatewayMoved) {
	grid_cell.AggregateMovedGateway(message)
	radar_beam.AggregateMovedGateway(message)
}

func ReprocessAntenna(antenna database.Antenna, movedTime time.Time) {
	grid_cell.ReprocessAntenna(antenna, movedTime)
	radar_beam.ReprocessAntenna(antenna, movedTime)
}
