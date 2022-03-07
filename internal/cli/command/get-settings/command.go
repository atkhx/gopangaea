package get_settings

import (
	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "settings"
	CliDescription = "показать детальные настройки пресета (wip, пока байты)"
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
