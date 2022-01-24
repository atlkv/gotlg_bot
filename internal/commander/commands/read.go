package commands

import common "tgbot/internal/commander/commands/common"


type Read struct {
	common.CMDModel
}

func (c *Read) Execute() error {
	return nil
}

func(c *Read) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewReadCommand() *Read {
	return &Read {
		common.CMDModel {
			CmdPath: "/read",
			Title: "Прочитать сущность",
		},
	}
}