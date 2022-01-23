package commands

import (
	"fmt"
)


type Delete struct {
	CMDModel
}

func (c *Delete) Execute() error {
	fmt.Print("Delete")
	return nil
}

func(c *Delete) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewDeleteCommand() *Delete {
	return &Delete {
		CMDModel {
			CmdPath: "/delete",
			Title: "Удалить сущность",
		},
	}
}