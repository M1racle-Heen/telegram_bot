package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5797499341:AAEyTUebxUuyDqnqIHrS27DiHjzoq5WK2jI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	// u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			j := update.Message.From.UserName + " said " + update.Message.Text
			ok := strings.Contains(j, "/exit")
			if ok {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "You have exited chat"))
				return
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, j)
			// msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}
