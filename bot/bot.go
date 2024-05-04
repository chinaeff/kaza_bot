package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"kaza_bot/config"
	"log"
)

var (
	botToken = config.BotToken
	botURL   = config.BotUrl
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal("error starting bot", err)
	}
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		HandleUpdate(bot, update)
	}
}
