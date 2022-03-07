package get_bank

import (
	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "bank"
	CliDescription = "показать номер выбранного банка и пресета"
)

type Device interface {
	ExecCommand(command device.Command) ([]byte, error)
}

type Command struct {
	device Device
}

func New(device Device) *Command {
	return &Command{device: device}
}

func (c *Command) Execute() (interface{}, error) {
	cmd := deviceCommand.New()

	r, err := c.device.ExecCommand(cmd)
	if err != nil {
		return nil, err
	}
	return cmd.ParseResponse(r)
}

func (c *Command) ParseArgs(args []string) error {
	return nil
}
