package owm

import (
	"The-weather-TGbot/internal/app"
	owm "github.com/briandowns/openweathermap"
	"log"
	"os"
)

type OWMAPI struct {
	app.API
}

func NewOWMAPI() *OWMAPI {
	owmAPI, exists := os.LookupEnv("owm_API_KEY")

	if !exists {
		log.Panic("Can't find OWM key in .env", exists)
	}
	return &OWMAPI{app.API{Key: owmAPI}}
}

func NewOWMData(key string) *app.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", key)
	if err != nil {
		log.Fatalln("OWM can't get data: ", err)
	}

	if app.WeatherData.GeoPos.Location == "" {
		w.CurrentByCoordinates(&owm.Coordinates{
			Longitude: app.WeatherData.GeoPos.Longitude,
			Latitude:  app.WeatherData.GeoPos.Latitude,
		})
	} else {
		w.CurrentByName(app.WeatherData.GeoPos.Location)
	}

	return &app.CurrentWeatherData{
		GeoPos: app.Coordinates{
			Longitude: w.GeoPos.Longitude,
			Latitude:  w.GeoPos.Latitude,
		},
		Main: app.Main{
			Temp:      w.Main.Temp,
			TempMin:   w.Main.TempMin,
			TempMax:   w.Main.TempMax,
			FeelsLike: w.Main.FeelsLike,
			Pressure:  w.Main.Pressure,
			Humidity:  w.Main.Humidity,
		},
	}
}
