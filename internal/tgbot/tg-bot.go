package tgbot

import (
	"The-weather-TGbot/internal/transport"
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

func (b *tgbot) getUpdates() *transport.CurrentWeatherData {
	var err error
	b.bot, err = tgbotapi.NewBotAPI(transport.Key)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	transport.WeatherData = &transport.CurrentWeatherData{GeoPos: transport.Coordinates{Longitude: 0, Latitude: 0, Location: ""}}

	for update := range updates {
		if update.Message != nil { // If we got a weather
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "" {
				transport.WeatherData.GeoPos.Longitude = update.Message.Location.Longitude
				transport.WeatherData.GeoPos.Latitude = update.Message.Location.Latitude
				transport.WeatherData.GeoPos.Location = ""
			} else {
				transport.WeatherData.GeoPos.Location = update.Message.Text
				transport.WeatherData.GeoPos.Longitude = 0
				transport.WeatherData.GeoPos.Latitude = 0
			}
		}
		b.ID = update.Message.Chat.ID
		b.MessageID = update.Message.MessageID
	}

	return transport.WeatherData
}

func (b *tgbot) sendMsg(answer string) {
	msg := tgbotapi.NewMessage(b.ID, answer)
	msg.ReplyToMessageID = b.MessageID
	b.bot.Send(msg)
}

type TGAPI struct {
	transport.API
}

func NewTGAPI() *TGAPI {
	tgAPI, exists := os.LookupEnv("tg_API_KEY")
	if !exists {
		log.Panic("Can't find TG-bot key in .env", exists)
	}
	return &TGAPI{transport.API{Key: tgAPI}}
}

func CreateTGBot(bot tgbot) *transport.CurrentWeatherData {
	return bot.getUpdates()
}

func CreateTGSender(bot tgbot) {
	bot.sendMsg(transport.Answer)
}
