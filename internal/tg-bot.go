package internal

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

/*type (
	BotAPI struct {
		Token           string     `json:"token"`
		Debug           bool       `json:"debug"`
		Buffer          int        `json:"buffer"`
		Self            User       `json:"-"`
		Client          HTTPClient `json:"-"`
		shutdownChannel chan interface{}
		apiEndpoint     string
	}
	User struct {
		ID                      int64  `json:"id"`
		IsBot                   bool   `json:"is_bot,omitempty"`
		FirstName               string `json:"first_name"`
		LastName                string `json:"last_name,omitempty"`
		UserName                string `json:"username,omitempty"`
		LanguageCode            string `json:"language_code,omitempty"`
		CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	}
	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
)*/

type tgbot struct {
	bot       *tgbotapi.BotAPI
	ID        int64
	MessageID int
}

func (b *tgbot) getUpdates() *CurrentWeatherData {
	var err error
	b.bot, err = tgbotapi.NewBotAPI(key)
	if err != nil {
		log.Panic("couldn't create a new BotAPI instance: ", err)
	}

	b.bot.Debug = true

	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	weatherData = &CurrentWeatherData{GeoPos: Coordinates{Longitude: 0, Latitude: 0, Location: ""}}

	for update := range updates {
		if update.Message != nil { // If we got a weather
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "" {
				weatherData.GeoPos.Longitude = update.Message.Location.Longitude
				weatherData.GeoPos.Latitude = update.Message.Location.Latitude
				weatherData.GeoPos.Location = ""
			} else {
				weatherData.GeoPos.Location = update.Message.Text
				weatherData.GeoPos.Longitude = 0
				weatherData.GeoPos.Latitude = 0
			}
		}
		b.ID = update.Message.Chat.ID
		b.MessageID = update.Message.MessageID
	}

	return weatherData
}

func (b *tgbot) sendMsg(answer string) {
	msg := tgbotapi.NewMessage(b.ID, answer)
	msg.ReplyToMessageID = b.MessageID
	b.bot.Send(msg)
}

type TGAPI struct {
	API
}

func NewTGAPI() *TGAPI {
	tgAPI, exists := os.LookupEnv("tg_API_KEY")
	if !exists {
		log.Panic("Can't find TG-bot key in .env", exists)
	}
	return &TGAPI{API{key: tgAPI}}
}

func createTGBot(bot tgbot) *CurrentWeatherData {
	return bot.getUpdates()
}
