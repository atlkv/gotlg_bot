package commands
import common "tgbot/internal/commander/commands/common"



type Delete struct {
	common.CMDModel
}

func (c *Delete) Execute() error {
	return nil
}

func(c *Delete) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewDeleteCommand() *Delete {
	return &Delete {
		common.CMDModel {
			CmdPath: "/delete",
			Title: "Удалить сущность",
		},
	}
}