package commands


type Update struct {
	CMDModel
}

func (c *Update) Execute() (string, error) {
	return "Update success", nil
}

func (c *Update) GetPath() string {
	return c.CmdPath
}

func(c *Update) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewUpdateCommand() *Update {
	return &Update {
		CMDModel {
			CmdPath: "/read",
			Title: "Read entity",
		},
	}
}