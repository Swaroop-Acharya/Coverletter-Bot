package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func setStringField(bot *tgbotapi.BotAPI, chatID int64, value, label string, setter func(string)) {
	if value == "" {
		bot.Send(tgbotapi.NewMessage(chatID, "❗ Please provide a value for "+label))
		return
	}

	setter(value)
	bot.Send(tgbotapi.NewMessage(chatID, "✅ "+label+" set to: "+value))
}
