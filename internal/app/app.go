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

type app struct {
	w   *owm.Owm
	bot *tgbot.TgBot
}

func getWeatherData(lon, lat float64, text string, weather *app) {
	if text == "" {
		weather.w.WeatherByCoord(lon, lat)
	} else {
		weather.w.WeatherByName(text)
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
	transport.APIKey = CreateAPIKey(tg)
	transport.Key = transport.GetAPIKey(transport.APIKey)
	bot := tgbot.StartTgBot()

	transport.APIKey = CreateAPIKey(openWeatherMap)
	transport.Key = transport.GetAPIKey(transport.APIKey)
	w := owm.StartOwm()

	application := app{bot: tgbot.NewTgBot(bot), w: owm.NewOwmApi(w)}

	lon, lat, text, messageID, ID := application.bot.ReadUpdates(application.bot.GetUpdates())

	getWeatherData(lon, lat, text, &application)

	answer := makeAnswerForMessanger(w.GeoPos.Longitude, w.GeoPos.Latitude, w.Main.Temp, w.Main.TempMin, w.Main.TempMax,
		w.Main.FeelsLike, w.Main.Pressure, w.Main.Humidity)

	application.bot.SendMessage(ID, answer, messageID)
}
