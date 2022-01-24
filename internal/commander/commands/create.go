package commands
import common "tgbot/internal/commander/commands/common"

type Create struct {
	common.CMDModel
}

func (c *Create) Execute() (string, error) {
	return "Create success", nil
}

func (c *Create) GetPath() string {
	return c.CmdPath
}

func(c *Create) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewCreateCommand() *Create {
	return &Create {
		common.CMDModel {
			CmdPath: "/create",
			Title: "Create entity",
		},
	}
}