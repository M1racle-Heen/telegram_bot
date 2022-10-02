package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/M1racle-Heen/bot/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{bot: bot, productService: productService}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message
		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)
		case "list":
			c.List(update.Message)
		default:
			c.Default(update.Message)
		}
	}
}

func String() string {
	return ""
}
