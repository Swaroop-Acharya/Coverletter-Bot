package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cld *coverletterData) setStringField(bot *tgbotapi.BotAPI, chatID int64, value string, field FieldKey) {
	if value == "" {
		bot.Send(tgbotapi.NewMessage(chatID, "❗ Please provide "+string(field)))
		return
	}
	switch field {
	case FieldTargetCompany:
		cld.TargetCompany = value
	case FieldCurrentCompany:
		cld.CurrentCompany = value
	case FieldTargetRole:
		cld.TargetRole = value
	case FieldCurrentRole:
		cld.CurrentRole = value
	case FieldUserName:
		cld.UserName = value
	case FieldUserEmail:
		cld.UserEmail = value
	case FieldUserPhone:
		cld.UserPhoneNumber = value
	case FieldIntro:
		cld.Intro = value
	case FieldClosing:
		cld.Closing = value
	}
	bot.Send(tgbotapi.NewMessage(chatID, "✅ "+ string(field)+" set to: "+value))
}
