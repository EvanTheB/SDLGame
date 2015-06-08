package tools

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetPlanets2() []*Body {
	data, err := ioutil.ReadFile("space/planets.json")

	var json_planets []*Body
	err = json.Unmarshal(data, &json_planets)
	if err != nil {
		log.Fatal(err)
	}

	return json_planets
}
