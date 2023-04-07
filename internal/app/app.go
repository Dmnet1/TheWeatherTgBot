package app

import (
	"The-weather-TGbot/internal/owm"
	"The-weather-TGbot/internal/tgbot"
	"fmt"
	"strconv"
)

const (
	openWeatherMap string = "openWeatherMap"
	tg             string = "tg"
)

var (
	WeatherData *CurrentWeatherData
	APIKey      KeyAPIGetter
	Key, Answer string
)

type (
	API struct {
		Key string
	}
	KeyAPIGetter interface {
		getKey() string
	}
)

func (a *API) getKey() string {
	return a.Key
}

func createAPIKey(resourseName string) KeyAPIGetter {
	var key KeyAPIGetter
	switch resourseName {
	case "tg":
		key = tgbot.NewTGAPI()
		return key
	case "openWeatherMap":
		key = owm.NewOWMAPI()
		return key
	}
	return key
}

func getAPIKey(key KeyAPIGetter) string {
	return key.getKey()
}

type (
	CurrentWeatherData struct {
		GeoPos Coordinates
		Main   Main
	}
	Coordinates struct {
		Longitude float64
		Latitude  float64
		Location  string
	}
	Main struct {
		Temp      float64
		TempMin   float64
		TempMax   float64
		FeelsLike float64
		Pressure  float64
		Humidity  int
	}
)

func (c *CurrentWeatherData) getWeatherParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) {
	return c.GeoPos.Longitude, c.GeoPos.Latitude, c.Main.Temp, c.Main.TempMin, c.Main.TempMax, c.Main.FeelsLike, c.Main.Pressure, c.Main.Humidity
}

func createDataWeather(apiName string) *CurrentWeatherData {
	switch apiName {
	case "openWeatherMap":
		WeatherData = owm.NewOWMData(Key)
		return WeatherData
	}
	return WeatherData
}

func createBot(apiName string) *CurrentWeatherData {
	switch apiName {
	case "tg":
		WeatherData = tgbot.CreateTGBot(tgbot.TgBot)
	}
	return WeatherData
}

func createSender(apiName string) {
	switch apiName {
	case "tg":
		tgbot.CreateTGSender(tgbot.TgBot)
	}
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
	Key = getAPIKey(APIKey)
	WeatherData = createBot(tg)
	APIKey = createAPIKey(openWeatherMap)
	Key = getAPIKey(APIKey)
	WeatherData = createDataWeather(openWeatherMap)
	Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity := WeatherData.getWeatherParam()
	Answer = makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity)
	createSender(tg)

}
