package owm

import (
	"The-weather-TGbot/internal/transport"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type OWMAPI struct {
	transport.API
}

func NewOWMAPI() *OWMAPI {
	owmAPI, exists := os.LookupEnv("owm_API_KEY")

	if !exists {
		log.Panic("Can't find OWM key in .env", exists)
	}
	return &OWMAPI{transport.API{Key: owmAPI}}
}

func NewOWMData(key string) *transport.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", key)
	if err != nil {
		log.Fatalln("OWM can't get data: ", err)
	}

	if transport.WeatherData.GeoPos.Location == "" {
		w.CurrentByCoordinates(&owm.Coordinates{
			Longitude: transport.WeatherData.GeoPos.Longitude,
			Latitude:  transport.WeatherData.GeoPos.Latitude,
		})
	} else {
		w.CurrentByName(transport.WeatherData.GeoPos.Location)
	}

	return &transport.CurrentWeatherData{
		GeoPos: transport.Coordinates{
			Longitude: w.GeoPos.Longitude,
			Latitude:  w.GeoPos.Latitude,
		},
		Main: transport.Main{
			Temp:      w.Main.Temp,
			TempMin:   w.Main.TempMin,
			TempMax:   w.Main.TempMax,
			FeelsLike: w.Main.FeelsLike,
			Pressure:  w.Main.Pressure,
			Humidity:  w.Main.Humidity,
		},
	}
}
