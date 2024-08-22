package state

import (
	"encoding/json"
	"log"
)

type State struct {
	State  string   `json:"state"`
	Cities []string `json:"cities"`
}

func (state State) GetCitiesAsJson() string {
	citiesJson, err := json.Marshal(state.Cities)
	if err != nil {
		log.Fatalf("Got error marshalling: %s", err)
	}
	return string(citiesJson)
}
