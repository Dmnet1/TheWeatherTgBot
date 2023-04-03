package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	openWeatherMap, exists := os.LookupEnv("owm_API_KEY")

	if exists {
		fmt.Println(openWeatherMap)
	}

	telegramBot, exists := os.LookupEnv("tg_API_KEY")

	if exists {
		fmt.Println(telegramBot)
	}

	w, err := owm.NewCurrent("C", "ru", "owm_API_KEY") // fahrenheit (imperial) with Russian output
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := tgbotapi.NewBotAPI("tg_API_KEY")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a dataWeather
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			w.CurrentByName(update.Message.Text)

			weatherByCountry := "Temp: " + fmt.Sprintf("%.2f\n", w.Main.Temp) + "Temp max: " +
				fmt.Sprintf("%.2f\n", w.Main.TempMax) + "Temp min: " + fmt.Sprintf("%.2f\n", w.Main.TempMin) +
				"Feels like: " + fmt.Sprintf("%.2f\n", w.Main.FeelsLike)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, weatherByCountry)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
