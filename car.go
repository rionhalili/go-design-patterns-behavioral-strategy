package main

type Car struct {
	timeSpent string
}

func (car *Car) GetToAirport() string {
	return `Driving to the Airport will take around ` + car.timeSpent + ` minutes`
}
