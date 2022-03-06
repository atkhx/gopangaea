package get_mode

import "github.com/atkhx/gopangaea/internal/cli/command"

const (
	GetMode    = "gm"
	CliCommand = "mode"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return GetMode
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
