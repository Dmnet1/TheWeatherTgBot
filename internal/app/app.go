package app

import (
	"The-weather-TGbot/internal/owm"
	"The-weather-TGbot/internal/tgbot"
	"The-weather-TGbot/internal/transport"
	"fmt"
	"strconv"
)

const (
	openWeatherMap string = "openWeatherMap"
	tg             string = "tg"
)

func makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) string {
	dataForMessanger := "Temp: " + fmt.Sprintf("%.2f\n", Temp) + "Temp max: " + fmt.Sprintf("%.2f\n", TempMax) +
		"Temp min: " + fmt.Sprintf("%.2f\n", TempMin) + "Feels like: " + fmt.Sprintf("%.2f\n", FeelsLike) +
		"Pressure: " + fmt.Sprintf("%.2f\n", Pressure) + "Humidity: " + strconv.Itoa(Humidity) +
		"Geo location: " + fmt.Sprintf("%.2f, %.2f\n", Latitude, Longitude)
	return dataForMessanger
}

func CreateDataWeather(apiName string) *transport.CurrentWeatherData {
	switch apiName {
	case "openWeatherMap":
		transport.WeatherData = owm.NewOWMData(transport.Key)
		return transport.WeatherData
	}
	return transport.WeatherData
}

func CreateBot(apiName string) *transport.CurrentWeatherData {
	switch apiName {
	case "tg":
		transport.WeatherData = tgbot.CreateTGBot(tgbot.TgBot)
	}
	return transport.WeatherData
}

func CreateSender(apiName string) {
	switch apiName {
	case "tg":
		tgbot.CreateTGSender(tgbot.TgBot)
	}
}

func CreateAPIKey(resourseName string) transport.KeyAPIGetter {
	var key transport.KeyAPIGetter
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

func Run() {
	transport.APIKey = CreateAPIKey(tg)
	transport.Key = transport.GetAPIKey(transport.APIKey)
	transport.WeatherData = CreateBot(tg)
	transport.APIKey = CreateAPIKey(openWeatherMap)
	transport.Key = transport.GetAPIKey(transport.APIKey)
	transport.WeatherData = CreateDataWeather(openWeatherMap)
	Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity := transport.WeatherData.GetWeatherParam()
	transport.Answer = makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure, Humidity)
	CreateSender(tg)

}
