package get_impulse_names

import "github.com/atkhx/gopangaea/internal/cli/command"

const (
	GetImpulseNames = "rns"
	CliCommand      = "impulse-names"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return GetImpulseNames
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
