package commands

import (
	"fmt"
)


type Read struct {
	CMDModel
}

func (c *Read) Execute() error {
	fmt.Print("Read")
	return nil
}

func(c *Read) Info() string {
	return c.CmdPath + " - " + c.Title
}

func NewReadCommand() *Read {
	return &Read {
		CMDModel {
			CmdPath: "/read",
			Title: "Прочитать сущность",
		},
	}
}