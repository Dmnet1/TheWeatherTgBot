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

type location interface {
	GetLoc() (float64, float64, string)
}

type mainWeatherParam interface {
	GetParam() (Longitude, Latitude, Temp, TempMin, TempMax, FeelsLike, Pressure float64, Humidity int)
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
	w     weather
	b     bot
	loc   location
	param mainWeatherParam
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

func getWeatherInfo(a app) {
	a.b.HandleUpdates()
	lon, lat, text := a.loc.GetLoc()
	if text == "" {
		a.w.WeatherByCoord(lon, lat)
	} else {
		a.w.WeatherByName(text)
	}
	longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure, humidity := a.param.GetParam()
	answer := makeAnswerForMessanger(longitude, latitude, temp, tempMin, tempMax, feelsLike, pressure, humidity)

	a.b.SendMsg(answer)
}

func Run() {
	b := tgbot.StartTgBot(makeApi(tg))
	w := owm.StartOwm(makeApi(openWeatherMap))
	application := app{b: tgbot.NewTgBot(b), w: owm.NewOwmApi(w)}

	/*lon, lat, text, messageID, ID := application.bot.ReadUpdates(application.bot.GetUpdates())
	getWeatherData(lon, lat, text, owm.NewOwmApi(w))
	answer := makeAnswerForMessanger(w.GeoPos.longitude, w.GeoPos.latitude, w.Main.temp, w.Main.tempMin, w.Main.tempMax,
		w.Main.feelsLike, w.Main.pressure, w.Main.humidity)
	application.bot.SendMessage(ID, answer, messageID)*/

	//it's realization with interface
	getWeatherInfo(application)
}

/*func getWeatherData(lon, lat float64, text string, weather weather) {
	if text == "" {
		weather.WeatherByCoord(lon, lat)
	} else {
		weather.WeatherByName(text)
	}
}*/
