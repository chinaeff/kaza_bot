package admin_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"kaza_bot/config"
	"log"
)

var (
	adminBotToken = config.AdminBotToken
	adminBotUrl   = config.AdminBotUrl
)

func StartAdminBot() {
	bot, err := tgbotapi.NewBotAPI(adminBotToken)
	if err != nil {
		log.Fatal("error starting admin bot", err)
	}

	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		HandleAdminUpdate(bot, update)
	}
}
