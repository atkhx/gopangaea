package get_impulse_name

import "github.com/atkhx/gopangaea/internal/cli/command"

const (
	GetImpulseName = "rn"
	CliCommand     = "impulse"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return GetImpulseName
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
