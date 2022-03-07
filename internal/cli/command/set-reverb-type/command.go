package set_reverb_type

import (
	"flag"

	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/set-reverb-type"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "set-reverb-type"
	CliDescription = "установить тип реверберации (0 - short, 1 - medium, 2 - long)"
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

	value := flagset.Int("value", 1, "value [0..2]")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	if err := deviceCommand.Validate(*value); err != nil {
		return err
	}

	c.value = *value
	return nil
}
