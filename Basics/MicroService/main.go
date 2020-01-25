package main

import (
	"encoding/json"
	"github.com/Rioverde/Go/MicroService/webportal"
	"log"
	"os"
)

type configuration struct {
	Webserver string `json:"webserver"`
}

func main() {
	file, err := os.Open("server.json")
	if err != nil {
		log.Fatal(err)
	}

	//we create new configuration
	config := new(configuration)
	//send it from json as decoded configuration structure
	json.NewDecoder(file).Decode(config)
	log.Println("Starting server on adress:", config.Webserver)

	webportal.RunWebPortal(config.Webserver)

}
