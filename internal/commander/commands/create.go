package commands


type Create struct {
	CMDModel
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
		CMDModel {
			CmdPath: "/create",
			Title: "Create entity",
		},
	}
}