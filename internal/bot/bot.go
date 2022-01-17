package bot

import (
	"tgbot/internal/commander"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)


type Bot struct {	
	tgbot	*tgbotapi.BotAPI	
}

func NewBot(token string) (*Bot, error) {
	tgbot, err := tgbotapi.NewBotAPI(token)
	
	if err != nil {
		return nil, err
	}

	return &Bot{tgbot: tgbot}, nil
}

func (b *Bot) Run(timeout int) {
	config := tgbotapi.NewUpdate(0)
	config.Timeout = timeout
	
	updates, _ := b.tgbot.GetUpdatesChan(config)
	
	commander := commander.NewCommander(b.tgbot)
	commander.Handle(updates)	
}

func (b *Bot) UserName() string {
	return b.tgbot.Self.UserName
}
