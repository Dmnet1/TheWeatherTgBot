package main

import (
	"The-weather-TGbot/internal"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	internal.Run()
	/*openWeatherMap := internal.GiveAPIKeyForOWM()

	telegramBot := internal.GiveAPIKeyForTGBot()

	w, err := owm.NewCurrent("C", "ru", openWeatherMap)
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := tgbotapi.NewBotAPI(telegramBot)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	var lon, lat = 0.0, 0.0

	for update := range updates {
		if update.Message != nil { // If we got a dataWeather
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "" {
				lon = update.Message.Location.Longitude
				lat = update.Message.Location.Latitude
				w.CurrentByCoordinates(&owm.Coordinates{
					Longitude: lon,
					Latitude:  lat,
				})
			} else {
				w.CurrentByName(update.Message.Text)
			}*/

	/*			weatherByCountry := "Temp: " + fmt.Sprintf("%.2f\n", w.Main.Temp) + "Temp max: " +
				fmt.Sprintf("%.2f\n", w.Main.TempMax) + "Temp min: " + fmt.Sprintf("%.2f\n", w.Main.TempMin) +
				"Feels like: " + fmt.Sprintf("%.2f\n", w.Main.FeelsLike) + "Geo location: " +
				fmt.Sprintf("%.2f, %.2f\n", w.GeoPos.Latitude, w.GeoPos.Longitude)*/

	/*	msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherByCountry)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}*/
}
