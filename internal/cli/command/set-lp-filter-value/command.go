package set_lp_filter_value

import (
	"flag"

	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/set-lp-filter-value"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "set-lp-filter-value"
	CliDescription = "установить значение LP-фильтра"
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

	// todo провалидировать min/max
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
