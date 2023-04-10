package app

import (
	"The-weather-TGbot/internal/owm"
	"The-weather-TGbot/internal/tgbot"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	openWeatherMap string = "openWeatherMap"
	tg             string = "tg"
)

type geoDataGetter interface {
	GetGeoData() (float64, float64, string)
}

type weatherParamGetter interface {
	GetWeatherParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int)
}

type bot interface {
	HandleUpdates() //приватный
	SendMsg(text string)
}

type weather interface {
	WeatherByCoord(longitude, latitude float64)
	WeatherByName(locationName string)
}

type app struct {
	w weather //w   *owm.Owm //тут должен быть интерфейс погоды
	b bot     //bot *tgbot.TgBot//тут бота
}

func getWeatherData(lon, lat float64, text string, weather weather) {
	if text == "" {
		weather.WeatherByCoord(lon, lat)
	} else {
		weather.WeatherByName(text)
	}
}

func makeApi(name string) string {
	var key string
	var exists bool
	switch name {
	case "tg":
		key, exists = os.LookupEnv("tg_API_KEY")
		if !exists {
			log.Panic("Can't find TG-bot key in .env", exists)
		}
	case "openWeatherMap":
		key, exists = os.LookupEnv("owm_API_KEY")
		if !exists {
			log.Panic("Can't find OWM key in .env", exists)
		}
	}
	return key
}

func makeAnswerForMessanger(Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int) string {
	dataForMessanger := "temp: " + fmt.Sprintf("%.2f\n", Temp) + "temp max: " + fmt.Sprintf("%.2f\n", TempMax) +
		"temp min: " + fmt.Sprintf("%.2f\n", TempMin) + "Feels like: " + fmt.Sprintf("%.2f\n", FeelsLike) +
		"pressure: " + fmt.Sprintf("%.2f\n", Pressure) + "humidity: " + strconv.Itoa(Humidity) +
		"Geo location: " + fmt.Sprintf("%.2f, %.2f\n", Latitude, Longitude)
	return dataForMessanger
}

func getWeatherInfo(b bot, w weather, d geoDataGetter, weather weatherParamGetter) {
	b.HandleUpdates()
	lon, lat, text := d.GetGeoData()
	if text == "" {
		w.WeatherByCoord(lon, lat)
	} else {
		w.WeatherByName(text)
	}
	longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure, humidity := weather.GetWeatherParam()
	answer := makeAnswerForMessanger(longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure, humidity)

	b.SendMsg(answer)
}

func Run() {
	b := tgbot.StartTgBot(makeApi(tg))
	w := owm.StartOwm(makeApi(openWeatherMap))

	//this is realization without interface
	application := app{b: tgbot.NewTgBot(b), w: owm.NewOwmApi(w)}
	/*lon, lat, text, messageID, ID := application.bot.ReadUpdates(application.bot.GetUpdates())
	getWeatherData(lon, lat, text, owm.NewOwmApi(w))
	answer := makeAnswerForMessanger(w.GeoPos.longitude, w.GeoPos.latitude, w.Main.temp, w.Main.tempMin, w.Main.tempMax,
		w.Main.feelsLike, w.Main.pressure, w.Main.humidity)
	application.bot.SendMessage(ID, answer, messageID)*/

	//it's realization with interface
	getWeatherInfo(tgbot.NewTgBot(bot), owm.NewOwmApi(w), tgbot.NewTgBot(bot), owm.NewOwmApi(w))

}
