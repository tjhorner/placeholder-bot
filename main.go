package main

import (
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("BOT_API_TOKEN")
	placeholder := os.Getenv("BOT_PLACEHOLDER_TEXT")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Chat.Type != "private" {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, placeholder)
		msg.ParseMode = tgbotapi.ModeMarkdown

		bot.Send(msg)
	}
}
