package internal

import (
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type OWMAPI struct {
	API
}

func NewOWMAPI() *OWMAPI {
	owmAPI, exists := os.LookupEnv("owm_API_KEY")

	if !exists {
		log.Panic("Can't find OWM key in .env", exists)
	}
	return &OWMAPI{API{key: owmAPI}}
}

func NewOWMData(key string) *CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", key)
	if err != nil {
		log.Fatalln("OWM can't get data: ", err)
	}

	if weatherData.GeoPos.Location == "" {
		w.CurrentByCoordinates(&owm.Coordinates{
			Longitude: weatherData.GeoPos.Longitude,
			Latitude:  weatherData.GeoPos.Latitude,
		})
	} else {
		w.CurrentByName(weatherData.GeoPos.Location)
	}

	return &CurrentWeatherData{
		GeoPos: Coordinates{
			Longitude: w.GeoPos.Longitude,
			Latitude:  w.GeoPos.Latitude,
		},
		Main: Main{
			Temp:      w.Main.Temp,
			TempMin:   w.Main.TempMin,
			TempMax:   w.Main.TempMax,
			FeelsLike: w.Main.FeelsLike,
			Pressure:  w.Main.Pressure,
			Humidity:  w.Main.Humidity,
		},
	}
}
