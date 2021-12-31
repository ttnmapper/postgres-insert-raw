package database

import "gorm.io/gorm/clause"

func GetRadarBeam(indexer RadarBeamIndexer) (RadarBeam, error) {
	var radarBeam RadarBeam
	radarBeam.AntennaID = indexer.AntennaID
	radarBeam.Level = indexer.Level
	radarBeam.Bearing = indexer.Bearing
	err := Db.FirstOrCreate(&radarBeam, &radarBeam).Error
	return radarBeam, err
}

func SaveRadarBeam(radarBeam RadarBeam) {
	Db.Save(&radarBeam)
}

func GetRadarBeamsForAntenna(antenna Antenna) []RadarBeam {
	var radarBeams []RadarBeam
	Db.Where("antenna_id = ?", antenna.ID).Find(&radarBeams)
	return radarBeams
}

func CreateRadarBeams(radarBeams []RadarBeam) error {
	// On conflict override
	tx := Db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&radarBeams)
	return tx.Error
}

func DeleteRadarBeamsForAntenna(antenna Antenna) {
	Db.Where(&RadarBeam{AntennaID: antenna.ID}).Delete(&RadarBeam{})
}
