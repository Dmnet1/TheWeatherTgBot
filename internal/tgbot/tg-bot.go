package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type tgBot struct {
	bot       *tgbotapi.BotAPI
	lon, Lat  float64
	text      string
	messageID int
	chatID    int64
}

func (t *tgBot) GetLoc() (float64, float64, string) {
	return t.lon, t.Lat, t.text
}

func (t *tgBot) HandleUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)
	for update := range updates {
		t.lon = update.Message.Location.Longitude
		t.Lat = update.Message.Location.Latitude
		t.text = update.Message.Text
		t.messageID = update.Message.MessageID
		t.chatID = update.Message.Chat.ID
	}
}

func (t *tgBot) SendMsg(answer string) {
	msg := tgbotapi.NewMessage(t.chatID, answer)
	msg.ReplyToMessageID = t.messageID
	t.bot.Send(msg)
}

func (t *tgBot) GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := t.bot.GetUpdatesChan(u)

	return updates
}

func NewTgBot(bot *tgbotapi.BotAPI) *tgBot {
	return &tgBot{
		bot:       bot,
		lon:       0,
		Lat:       0,
		text:      "",
		messageID: 0,
		chatID:    0,
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

/*func (t *tgBot) ReadUpdates(updates tgbotapi.UpdatesChannel) (lon float64, lat float64, text string, messageID int, ID int64) {
	for update := range updates {
		lon = update.Message.Location.Longitude
		lat = update.Message.Location.Latitude
		text = update.Message.text
		messageID = update.Message.messageID
		ID = update.Message.Chat.ID

		return lon, lat, text, messageID, ID
	}
	return lon, lat, text, messageID, ID
}*/

/*func (t *tgBot) SendMessage(ID int64, answer string, messageID int) {
	msg := tgbotapi.NewMessage(ID, answer)
	msg.ReplyToMessageID = messageID
	t.bot.Send(msg)
}*/
