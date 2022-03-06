package get_device

import "github.com/atkhx/gopangaea/internal/cli/command"

const (
	GetDevice  = "amtdev"
	CliCommand = "device"
)

type Command struct{}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return GetDevice
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
