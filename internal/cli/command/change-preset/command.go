package get_bank

import (
	"flag"

	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/change-preset"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "change-preset"
	CliDescription = "переключить выбранный пресет"
)

type Device interface {
	ExecCommand(command device.Command) ([]byte, error)
}

type Command struct {
	device Device
	preset int
	bank   int
}

func New(device Device) *Command {
	return &Command{device: device}
}

func (c *Command) Execute() (interface{}, error) {
	cmd, err := deviceCommand.NewWithArgs(c.bank, c.preset)
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
	flagset := flag.NewFlagSet("change-preset", flag.ContinueOnError)

	bank := flagset.Int("bank", 0, "index of bank [0..9]")
	preset := flagset.Int("preset", 0, "index of preset [0..9]")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	if err := deviceCommand.Validate(*bank, *preset); err != nil {
		return err
	}

	c.preset = *preset
	c.bank = *bank

	return nil
}
