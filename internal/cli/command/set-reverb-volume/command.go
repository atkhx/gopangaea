package set_reverb_volume

import (
	"flag"

	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/set-reverb-volume"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "set-reverb-volume"
	CliDescription = "установить громкость реверберации"
)

type Device interface {
	ExecCommand(command device.Command) ([]byte, error)
}

type Command struct {
	device Device
	value  int
}

func New(device Device) *Command {
	return &Command{device: device}
}

func (c *Command) Execute() (interface{}, error) {
	cmd, err := deviceCommand.NewWithArgs(c.value)
	if err != nil {
		return nil, err
	}
	r, err := c.device.ExecCommand(cmd)
	if err != nil {
		return nil, err
	}
	return cmd.ParseResponse(r)
}

func (c *Command) ParseArgs(args []string) error {
	flagset := flag.NewFlagSet(CliCommand, flag.ContinueOnError)

	value := flagset.Int("value", 16, "value [0..31]")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	if err := deviceCommand.Validate(*value); err != nil {
		return err
	}

	c.value = *value

	return nil
}
