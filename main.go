package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

type coverletterData struct {
	UserName        string
	UserPhoneNumber string
	UserEmail       string
	Skills          []string
	TargetCompany   string
	TargetRole      string
	CurrentCompany  string
	CurrentRole     string
	Points          []string
	Intro           string
	Closing         string
}

func main() {
	err := godotenv.Load()
	botAPIToken := os.Getenv("API_TOKEN")
	if botAPIToken == "" {
		log.Fatal("API_TOKEN is not set")
	}
	
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
			cmdArgs := update.Message.CommandArguments()
			switch update.Message.Command() {
			case "start":
				cld.handleStart(bot, chatID)
			case "help":
				cld.handleHelp(bot, chatID)
			case "username":
				cld.setStringField(bot, chatID, cmdArgs, FieldUserName)
			case "useremail":
				cld.setStringField(bot, chatID, cmdArgs, FieldUserEmail)
			case "userphonenumber":
				cld.setStringField(bot, chatID, cmdArgs, FieldUserPhone)
			case "targetcompany":
				cld.setStringField(bot, chatID, cmdArgs, FieldTargetCompany)
			case "currentcompany":
				cld.setStringField(bot, chatID, cmdArgs, FieldCurrentCompany)
			case "targetrole":
				cld.setStringField(bot, chatID, cmdArgs, FieldTargetRole)
			case "currentrole":
				cld.setStringField(bot, chatID, cmdArgs, FieldCurrentRole)
			case "skills":
				cld.setSkills(bot, chatID, cmdArgs)
			case "intro":
				cld.setStringField(bot, chatID, cmdArgs, FieldIntro)
			case "closing":
				cld.setStringField(bot, chatID, cmdArgs, FieldClosing)
			case "preview":
				cld.handlePreview(bot, chatID)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Command unavailable")
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}
		}
	}
}
