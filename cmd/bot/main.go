package main

import (
	"log"
	"os"

	"github.com/M1racle-Heen/bot/internal/app/commands"
	"github.com/M1racle-Heen/bot/internal/service/product"
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

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()
	command := commands.NewCommander(bot, productService)
	for update := range updates {
		command.HandleUpdate(update)
	}
}
