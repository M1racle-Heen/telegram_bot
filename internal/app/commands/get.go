package commands

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("Wrong argument: ", err)
		return
	}

	product, err := c.productService.Get(idx - 1)
	if err != nil {
		log.Printf("fail to get product with index: %d and got error %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}
