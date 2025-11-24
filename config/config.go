package config

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func InitBot() {
	var err error
	token := os.Getenv("BOT_TOKEN")
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
}
