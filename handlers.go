package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (cld *coverletterData) handleStart(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "👋 Welcome to SmartCover Bot!\nUse /help to see available commands.")
	bot.Send(msg)
}

func (cld *coverletterData) handleHelp(bot *tgbotapi.BotAPI, chatID int64) {
	helpText := `📖 SmartCover Bot Commands
/start - Start the bot
/help - Show this help menu
/targetcompany - Set target company
/currentcompany - Set your current company (optional)
/role - Set job title
/recruiter - Set recruiter or hiring manager
/name - Set your nam
/email - Set your email
/phone - Set phone number
/skills - Set skills summary (use ; as separator)
/intro - Set intro paragraph
/closing - Set closing paragraph
/points - Set bullet points (use ; as separator)
/location - Set location (optional)
/preview - Preview details
/printpdf - Generate and send the PDF`
	msg := tgbotapi.NewMessage(chatID, helpText)
	bot.Send(msg)
}
