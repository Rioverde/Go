package main

import (
	"fmt"
	"log"
	"os"

	// Shortening the import reference name seems to make it a bit easier
	owm "github.com/briandowns/openweathermap"
)

func main() {

	var apiKey = os.Getenv("OWM_API_KEY")
	w, err := owm.NewCurrent("K", "EN", apiKey) // (internal - OpenWeatherMap reference for kelvin) with English output
	if err != nil {
		log.Fatalln(err)
	}

	w.CurrentByName("Phoenix,AZ")
	fmt.Println(w)

}
