package tgbot

import (
	"The-weather-TGbot/internal/transport"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

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

type TgBot struct {
	Bot *tgbotapi.BotAPI
}

func NewTgBot(bot *tgbotapi.BotAPI) *TgBot {
	return &TgBot{Bot: bot}
}

func StartTgBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(transport.Key)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func (t *TgBot) GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.Bot.GetUpdatesChan(u)

	return updates
}

func (t *TgBot) ReadUpdates(updates tgbotapi.UpdatesChannel) (lon float64, lat float64, text string, messageID int, ID int64) {
	for update := range updates {
		lon = update.Message.Location.Longitude
		lat = update.Message.Location.Latitude
		text = update.Message.Text
		messageID = update.Message.MessageID
		ID = update.Message.Chat.ID

		return lon, lat, text, messageID, ID
	}
	return lon, lat, text, messageID, ID
}

func (t *TgBot) SendMessage(ID int64, answer string, messageID int) {
	msg := tgbotapi.NewMessage(ID, answer)
	msg.ReplyToMessageID = messageID
	t.Bot.Send(msg)
}
