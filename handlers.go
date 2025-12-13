package main

import (
	"strings"

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
	bot.Send(tgbotapi.NewMessage(chatID, "✅ "+string(field)+" set to: "+value))
}

func (cld *coverletterData) setSkills(bot *tgbotapi.BotAPI, chatID int64, cmdArgs string) {
	if cmdArgs == "" {
		bot.Send(tgbotapi.NewMessage(chatID, "❗ Please provide skills separated by `;`\nExample:\n/skills Go; Docker; AWS"))
		return
	}

	skills := parseList(cmdArgs)
	if len(skills) == 0 {
		bot.Send(tgbotapi.NewMessage(
			chatID,
			"❗ No valid skills found. Use `;` to separate skills.",
		))
		return
	}

	cld.Skills = skills
	msg := tgbotapi.NewMessage(
		chatID,
		"✅ Skills updated:\n• "+strings.Join(skills, "\n• "),
	)
	bot.Send(msg)
}

func (cld *coverletterData) handlePreview(bot *tgbotapi.BotAPI, chatID int64) {
	if cld.isEmpty() {
		bot.Send(tgbotapi.NewMessage(
			chatID,
			"❗ No details found yet.\nStart by using /targetcompany, /role, etc.",
		))
		return
	}

	preview := cld.formatPreview()

	msg := tgbotapi.NewMessage(chatID, preview)
	msg.ParseMode = "Markdown"
	bot.Send(msg)
}
