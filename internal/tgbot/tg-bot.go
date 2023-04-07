package tgbot

import (
	"The-weather-TGbot/internal/app"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

type tgbot struct {
	bot       *tgbotapi.BotAPI
	ID        int64
	MessageID int
}

var TgBot tgbot

func (b *tgbot) getUpdates() *app.CurrentWeatherData {
	var err error
	b.bot, err = tgbotapi.NewBotAPI(app.Key)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	app.WeatherData = &app.CurrentWeatherData{GeoPos: app.Coordinates{Longitude: 0, Latitude: 0, Location: ""}}

	for update := range updates {
		if update.Message != nil { // If we got a weather
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "" {
				app.WeatherData.GeoPos.Longitude = update.Message.Location.Longitude
				app.WeatherData.GeoPos.Latitude = update.Message.Location.Latitude
				app.WeatherData.GeoPos.Location = ""
			} else {
				app.WeatherData.GeoPos.Location = update.Message.Text
				app.WeatherData.GeoPos.Longitude = 0
				app.WeatherData.GeoPos.Latitude = 0
			}
		}
		b.ID = update.Message.Chat.ID
		b.MessageID = update.Message.MessageID
	}

	return app.WeatherData
}

func (b *tgbot) sendMsg(answer string) {
	msg := tgbotapi.NewMessage(b.ID, answer)
	msg.ReplyToMessageID = b.MessageID
	b.bot.Send(msg)
}

type TGAPI struct {
	app.API
}

func NewTGAPI() *TGAPI {
	tgAPI, exists := os.LookupEnv("tg_API_KEY")
	if !exists {
		log.Panic("Can't find TG-bot key in .env", exists)
	}
	return &TGAPI{app.API{Key: tgAPI}}
}

func CreateTGBot(bot tgbot) *app.CurrentWeatherData {
	return bot.getUpdates()
}

func CreateTGSender(bot tgbot) {
	bot.sendMsg(app.Answer)
}
