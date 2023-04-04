package internal

import (
	"log"
	"os"
)

func GiveAPIKeyForTGBot() string {
	telegramBot, exists := os.LookupEnv("tg_API_KEY")

	if !exists {
		log.Panic("Can't find TG-bot key in .env", exists)
	}
	return telegramBot
}
