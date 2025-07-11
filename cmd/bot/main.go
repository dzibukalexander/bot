package main

import (
	"log"
	"os"

	"github.com/dzibukalexander/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil { // If we don't got a message
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehaviour(bot, update.Message)
		}
	}
}

func listCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message, productService *product.Service) {
	titleList := "Products:\n\n"
	for _, product := range productService.ListProduct() {
		titleList += product.Title + "\n"
	}

	newMessage := tgbotapi.NewMessage(msg.Chat.ID, titleList)
	bot.Send(newMessage)
}

func helpCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "/help - help"))
}

func defaultBehaviour(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	newMessage := tgbotapi.NewMessage(msg.Chat.ID, msg.Text)
	newMessage.ReplyToMessageID = msg.MessageID
	bot.Send(newMessage)
}
