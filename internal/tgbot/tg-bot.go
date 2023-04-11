package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	Bot       *tgbotapi.BotAPI
	Lon, Lat  float64
	Text      string
	MessageID int
	ChatID    int64
}

func (t *TgBot) GetLoc() (float64, float64, string) {
	return t.Lon, t.Lat, t.Text
}

func (t *TgBot) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.Bot.GetUpdatesChan(u)
	for update := range updates {
		t.Lon = update.Message.Location.Longitude
		t.Lat = update.Message.Location.Latitude
		t.Text = update.Message.Text
		t.MessageID = update.Message.MessageID
		t.ChatID = update.Message.Chat.ID
	}
}

func (t *TgBot) SendMsg(answer string) {
	msg := tgbotapi.NewMessage(t.ChatID, answer)
	msg.ReplyToMessageID = t.MessageID
	t.Bot.Send(msg)
}

func (t *TgBot) GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.Bot.GetUpdatesChan(u)

	return updates
}

func NewTgBot(bot *tgbotapi.BotAPI) *TgBot {
	return &TgBot{
		Bot:       bot,
		Lon:       0,
		Lat:       0,
		Text:      "",
		MessageID: 0,
		ChatID:    0,
	}
}

func StartTgBot(key string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

/*func (t *TgBot) ReadUpdates(updates tgbotapi.UpdatesChannel) (lon float64, lat float64, text string, messageID int, ID int64) {
	for update := range updates {
		lon = update.Message.Location.Longitude
		lat = update.Message.Location.Latitude
		text = update.Message.Text
		messageID = update.Message.MessageID
		ID = update.Message.Chat.ID

		return lon, lat, text, messageID, ID
	}
	return lon, lat, text, messageID, ID
}*/

/*func (t *TgBot) SendMessage(ID int64, answer string, messageID int) {
	msg := tgbotapi.NewMessage(ID, answer)
	msg.ReplyToMessageID = messageID
	t.Bot.Send(msg)
}*/
