package get_impulse_name

import (
	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-name"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "impulse"
	CliDescription = "показать название выбранного импульса"
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
