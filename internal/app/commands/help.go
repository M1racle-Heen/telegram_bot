package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	var (
		help = "help"
		list = "list"
		get  = "get"
	)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("/%-10[1]s - %[1]s\n/%-13[2]s - %[2]s products\n/%-11[3]s - %[3]s some value", help, list, get),
	)
	c.bot.Send(msg)
}
