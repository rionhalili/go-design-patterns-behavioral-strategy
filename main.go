package main

import "fmt"

func main() {
	fmt.Println("====================================")

	transportationContext := &TransportationContext{}

	transportationContext.setStrategy(&Walking{timeSpent: "300"})
	fmt.Println(transportationContext.execute())

	transportationContext.setStrategy(&Car{timeSpent: "30"})
	fmt.Println(transportationContext.execute())
}
