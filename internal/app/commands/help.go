package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	var (
		help = "help"
		list = "list"
	)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		fmt.Sprintf("/%-10[1]s - %[1]s\n/%-13[2]s - %[2]s products", help, list),
	)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["help"] = (*Commander).Help
}
