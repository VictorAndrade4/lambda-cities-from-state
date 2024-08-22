package main

import (
	"fmt"
	"lambda-cities-from-state/state"
)

func main() {
	citiesAsJson := state.FromAlias("MG").GetCitiesAsJson()
	fmt.Println(citiesAsJson)
}
