package get_bank

import (
	"github.com/atkhx/gopangaea/internal/cli/command"
)

const (
	GetBank    = "gb"
	CliCommand = "bank"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return GetBank
}

func (c Command) GetResponseLength() int {
	return ResponseLength
}

func (c Command) Config() command.Config {
	return command.Config{
		Command: CliCommand,
	}
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}

func (c *Command) ParseArgs(args []string) error {
	return nil
}
