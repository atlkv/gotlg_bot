package commander

import (
	"fmt"
	cmd "tgbot/internal/commander/commands"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Commander struct {
	tgbot	*tgbotapi.BotAPI
	commands map[string]cmd.CMDExecuter
}

func NewCommander(b *tgbotapi.BotAPI) *Commander{
	create := cmd.NewCreateCommand()
	update := cmd.NewUpdateCommand()
	
	return &Commander{
		tgbot: b,
		commands: map[string]cmd.CMDExecuter{
			create.GetPath(): create,
			update.GetPath(): update,
		},
	}
}

func(c *Commander) Handle(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			c.message(update.Message)
		}
	}
}

func (c *Commander) message(message *tgbotapi.Message) {
	defer recoveryFunction()
	var replyMsg string

	if !message.IsCommand() {
		replyMsg = 
			"You message is not command.\nPlease see available commands: \n\n" + 
			c.getAvailableCommands()

		c.tgbot.Send(c.makeReply(message.Chat.ID, replyMsg))
		return
	}

	cmd, ok := c.GetCommand(message.Text)
	
	if !ok {
		replyMsg = 
			"Yor command: " + message.Text + "not exist \nPlease see available commands: \n\n" +
			c.getAvailableCommands()

		c.tgbot.Send(c.makeReply(message.Chat.ID, replyMsg))
		return
	}

	_, err := cmd.Execute()
	
	if err != nil {
		replyMsg = 
			"Unsuccess command: " + cmd.Info() + "\n\n" + 
			"Error: " + err.Error()

		c.tgbot.Send(c.makeReply(message.Chat.ID, replyMsg))
		return
	}

	replyMsg = "Success command: " + cmd.Info()

	c.tgbot.Send(c.makeReply(message.Chat.ID, replyMsg))
}

func recoveryFunction() {
	if recoveryMessage := recover(); recoveryMessage != nil {
		fmt.Println(recoveryMessage)
	}
	fmt.Println("This is recovery function...")
}

func (c *Commander) makeReply(chatID int64, message string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(
		chatID,
		message)
}

func (c *Commander) getAvailableCommands() string {	
	var info string

	for _, cmd := range c.commands {
		info = info + cmd.Info() + "\n"	
	}
	
	return info
}

func (c *Commander) GetCommand(command string) (cmd.CMDExecuter, bool) {
	cmd, err := c.commands[command] 
	
	return cmd, err
}
