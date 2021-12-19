package database

import "time"

type GridCell struct {
	ID        uint
	AntennaID uint `gorm:"UNIQUE_INDEX:idx_grid_cell"`

	X int `gorm:"UNIQUE_INDEX:idx_grid_cell"`
	Y int `gorm:"UNIQUE_INDEX:idx_grid_cell"`
	// Z is always 19

	LastUpdated time.Time

	BucketHigh     uint32
	Bucket100      uint32
	Bucket105      uint32
	Bucket110      uint32
	Bucket115      uint32
	Bucket120      uint32
	Bucket125      uint32
	Bucket130      uint32
	Bucket135      uint32
	Bucket140      uint32
	Bucket145      uint32
	BucketLow      uint32
	BucketNoSignal uint32
}

type GridCellIndexer struct {
	AntennaID uint
	X         int
	Y         int
}

type RadarBeam struct {
	ID uint

	AntennaID uint `gorm:"UNIQUE_INDEX:idx_radar_beam"`
	Level     int  `gorm:"UNIQUE_INDEX:idx_radar_beam"`
	Bearing   uint `gorm:"UNIQUE_INDEX:idx_radar_beam"`

	Samples     uint
	LastUpdated time.Time

	DistanceMax float64
	Distance2nd float64
}

type RadarBeamIndexer struct {
	AntennaID uint
	Level     int
	Bearing   uint
}
