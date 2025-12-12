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

	users := map[int64]*coverletterData{}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			chatID := update.Message.Chat.ID
			if users[chatID] == nil {
				users[chatID] = &coverletterData{}
			}
			cld := users[chatID]
			switch update.Message.Command() {
			case "start":
				cld.handleStart(bot, chatID)
			case "help":
				cld.handleHelp(bot, chatID)
			case "targetcompany":
				cld.setTargetCompany(bot,chatID,update.Message.CommandArguments())
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command unavailable")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
