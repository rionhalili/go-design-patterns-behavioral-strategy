package main

type TransportationContext struct {
	strategy TransportationStrategy
}

func (context *TransportationContext) setStrategy(strategy TransportationStrategy) {
	context.strategy = strategy
}

func (context *TransportationContext) execute() string {
	return context.strategy.GetToAirport()
}
