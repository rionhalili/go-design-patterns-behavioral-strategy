package main

type Walking struct {
	timeSpent string
}

func (walking *Walking) GetToAirport() string {
	return `Walking to the Airport will take around ` + walking.timeSpent + ` minutes`
}
