package set_hp_filter_state

import (
	"flag"

	deviceCommand "github.com/atkhx/gopangaea/internal/pkg/commands/set-hp-filter-state"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

const (
	CliCommand     = "set-hp-filter-state"
	CliDescription = "включить / выключить HP-фильтр"
)

type Device interface {
	ExecCommand(command device.Command) ([]byte, error)
}

type Command struct {
	device Device
	value  bool
}

func New(device Device) *Command {
	return &Command{device: device}
}

func (c *Command) Execute() (interface{}, error) {
	cmd := deviceCommand.NewWithArgs(c.value)
	r, err := c.device.ExecCommand(cmd)
	if err != nil {
		return nil, err
	}
	return cmd.ParseResponse(r)
}

func (c *Command) ParseArgs(args []string) error {
	flagset := flag.NewFlagSet(CliCommand, flag.ContinueOnError)

	value := flagset.Bool("value", false, "")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	c.value = *value
	return nil
}