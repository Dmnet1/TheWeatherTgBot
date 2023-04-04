package internal

import (
	"log"
	"os"
)

func GiveAPIKeyForOWM() string {
	openWeatherMap, exists := os.LookupEnv("owm_API_KEY")

	if !exists {
		log.Panic("Can't find OWM key in .env", exists)
	}
	return openWeatherMap
}
