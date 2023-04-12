package app

import (
	"The-weather-TGbot/internal/owm"
	"The-weather-TGbot/internal/tgbot"
	"fmt"
	"strconv"
)

type location interface {
	GetLoc() (float64, float64, string)
}

type mainWeatherParam interface {
	GetParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int)
}

type bot interface {
	HandleUpdates()
	SendMsg(text string)
}

type weather interface {
	WeatherByCoord(longitude, latitude float64)
	WeatherByName(locationName string)
}

type key interface {
	GetKey() string
}

type app struct {
	w     weather
	b     bot
	loc   location
	param mainWeatherParam
}

func makeAnswerForMessanger(a app) string {
	longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure, humidity := a.param.GetParam()
	return fmt.Sprintf("temp: %.2f\ntemp max: %.2f\ntemp min: %.2f\nFeels like: %.2f\npressure: %.2f\nhumidity: %s\nGeo location: %.2f, %.2f\n",
		temp, tempMax, tempMin, feelsLike, pressure, strconv.Itoa(humidity), latitude, longitude)
}

func getWeatherInfo(a app) {
	a.b.HandleUpdates()
	lon, lat, text := a.loc.GetLoc()
	if text == "" {
		a.w.WeatherByCoord(lon, lat)
	} else {
		a.w.WeatherByName(text)
	}
	answer := makeAnswerForMessanger(a)

	a.b.SendMsg(answer)
}

func getKey(k key) string {
	return k.GetKey()
}

func Run() {
	b := tgbot.StartTgBot(getKey(tgbot.NewKey()))
	w := owm.StartOwm(getKey(owm.NewKey()))
	application := app{b: tgbot.NewTgBot(b), w: owm.NewOwmApi(w)}
	getWeatherInfo(application)
}
