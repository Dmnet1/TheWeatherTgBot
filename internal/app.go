package internal

import (
	"fmt"
	"strconv"
)

const openWeatherMap string = "openWeatherMap"
const tg string = "tg"

var weather Weather

var tgBot tgbot

var bot Tranporter

var weatherData *CurrentWeatherData

var APIKey KeyAPIGetter

var key string

type API struct {
	key string
}

type KeyAPIGetter interface {
	getKey() string
}

func (a *API) getKey() string {
	return a.key
}

func createAPIKey(resourseName string) KeyAPIGetter {
	var key KeyAPIGetter
	switch resourseName {
	case "tg":
		key = NewTGAPI()
		return key
	case "openWeatherMap":
		key = NewOWMAPI()
		return key
	}
	return key
}

func getAPIKey(key KeyAPIGetter) string {
	return key.getKey()
}

type CurrentWeatherData struct {
	GeoPos Coordinates
	Main   Main
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
	Location  string
}

type Main struct {
	Temp      float64
	TempMin   float64
	TempMax   float64
	FeelsLike float64
	Pressure  float64
	Humidity  int
}

type Tranporter interface {
	getLocation() string
	getGeoPos() (float64, float64)
}

func (c *CurrentWeatherData) getLocation() string {
	return c.GeoPos.Location
}

func (c *CurrentWeatherData) getGeoPos() (float64, float64) {
	return c.GeoPos.Longitude, c.GeoPos.Latitude
}

type Weather interface {
	getAllOfWeather() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int)
}

func (c *CurrentWeatherData) getAllOfWeather() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) {
	return c.GeoPos.Longitude, c.GeoPos.Latitude, c.Main.Temp, c.Main.TempMin, c.Main.TempMax, c.Main.FeelsLike, c.Main.Pressure, c.Main.Humidity
}

func getWeatherParam(w Weather) (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) {
	return w.getAllOfWeather()
}

func createDataWeather(resourseName string) Weather {
	var data Weather
	switch resourseName {
	case "openWeatherMap":
		data = NewOWMData(key)
		return data
	}
	return data
}

func createBot(resourseName string) *CurrentWeatherData {
	var bot *CurrentWeatherData
	switch resourseName {
	case "tg":
		bot = createTGBot(tgBot)
	}

	return bot
}

func makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) string {
	dataForMessanger := "Temp: " + fmt.Sprintf("%.2f\n", Temp) + "Temp max: " + fmt.Sprintf("%.2f\n", TempMax) +
		"Temp min: " + fmt.Sprintf("%.2f\n", TempMin) + "Feels like: " + fmt.Sprintf("%.2f\n", FeelsLike) +
		"Pressure: " + fmt.Sprintf("%.2f\n", Pressure) + "Humidity: " + strconv.Itoa(Humidity) +
		"Geo location: " + fmt.Sprintf("%.2f, %.2f\n", Latitude, Longitude)
	return dataForMessanger
}

func Run() {
	APIKey = createAPIKey(tg)
	key = getAPIKey(APIKey)
	bot = createBot(tg)

	APIKey = createAPIKey(openWeatherMap)
	key = getAPIKey(APIKey)
	weather = createDataWeather(openWeatherMap)

	Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity := getWeatherParam(weather)
	answer := makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity)
	tgBot.sendMsg(answer)

	//тут должна быть код, где используется key для tg
}
