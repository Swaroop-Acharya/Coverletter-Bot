package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type coverletterData struct {
	UserName        string
	UserPhoneNumber int
	UserEmail       string
	Skills          []string
	TargetCompany   string
	TargetRole      string
	CurrentCompany  string
	CurrentRole     string
	Location        string
	Points          []string
	Intro           string
	Closing         string
}

func handleHelp(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
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
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpText)
	bot.Send(msg)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botAPIToken := os.Getenv("API_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botAPIToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Autherized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "👋 Welcome to SmartCover Bot!\nUse /help to see available commands.")
				bot.Send(msg)
			case "help":
				handleHelp(bot, &update)
			case "targetcompany":

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command unavailable")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
