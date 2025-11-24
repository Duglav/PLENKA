package main

import (
	"log"
	"plenkabot/config"
	"plenkabot/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Bot.Token)
	if err != nil {
		log.Panic("Ошибка запуска бота:", err)
	}

	config.Bot = bot

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		handlers.HandleUpdate(update)
	}
}
