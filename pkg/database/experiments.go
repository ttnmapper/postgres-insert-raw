package database

func GetExperimentList() []Experiment {
	var experiments []Experiment
	Db.Order("name asc").Find(&experiments)
	return experiments
}

func FindExperiment(experiment string) []Experiment {
	var experiments []Experiment
	experiment = "%" + experiment + "%"
	Db.Where("name LIKE ?", experiment).Order("name asc").Find(&experiments)
	return experiments
}
