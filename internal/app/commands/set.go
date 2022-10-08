package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Set(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	product, err := c.productService.Set(args)
	if err != nil {
		log.Printf("fail to add product: %s with error %v", product, err)
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title+" was added")
	c.bot.Send(msg)
}
