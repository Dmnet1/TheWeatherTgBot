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

type Owm struct {
	W *owm.CurrentWeatherData
}

func NewOwmApi(w *owm.CurrentWeatherData) *Owm {
	return &Owm{W: w}
}

func StartOwm() *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", transport.Key)
	if err != nil {
		log.Fatalln("OWM can't get data: ", err)
	}

	return w
}

func (o *Owm) WeatherByCoord(longitude, latitude float64) {
	o.W.CurrentByCoordinates(&owm.Coordinates{
		Longitude: longitude,
		Latitude:  latitude,
	})
}

func (o *Owm) WeatherByName(locationName string) {
	o.W.CurrentByName(locationName)
}
