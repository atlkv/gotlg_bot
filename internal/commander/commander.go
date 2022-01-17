package commander

import (
	//"log"
	"fmt"
	//"runtime/debug"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	tgbot	*tgbotapi.BotAPI
}

func NewCommander(b *tgbotapi.BotAPI) *Commander{
	return &Commander{
		tgbot: b,
	}
}

func(c *Commander) Handle(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {	
			c.Message(update.Message)
		}
	}
	fmt.Println("The function executes Completely")
}

func (c *Commander) Message(message *tgbotapi.Message) {
	defer recoveryFunction()

	if !message.IsCommand() {
		rspMessage := tgbotapi.NewMessage(message.Chat.ID, "Your message is not command!")
		c.tgbot.Send(rspMessage)
		return
	}
	
	rspMessage := tgbotapi.NewMessage(message.Chat.ID, "Success handle command: " + message.Text)
	c.tgbot.Send(rspMessage)
}

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}
