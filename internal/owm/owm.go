package owm

import (
	owm "github.com/briandowns/openweathermap"
	"log"
)

type Owm struct {
	W                                                                *owm.CurrentWeatherData
	longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure float64
	humidity                                                         int
}

func (o *Owm) GetWeatherParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) {
	return o.longitude, o.latitude, o.temp, o.tempMin, o.tempMax, o.feelsLike, o.pressure, o.humidity
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

func NewOwmApi(w *owm.CurrentWeatherData) *Owm {
	return &Owm{
		W:         w,
		longitude: 0,
		latitude:  0,
		temp:      0,
		tempMin:   0,
		tempMax:   0,
		feelsLike: 0,
		pressure:  0,
		humidity:  0,
	}
}

func StartOwm(key string) *owm.CurrentWeatherData {
	w, err := owm.NewCurrent("C", "ru", key)
	if err != nil {
		log.Fatalln("OWM can't get data: ", err)
	}
	return w
}
