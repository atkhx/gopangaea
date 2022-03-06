package change_preset

import (
	"errors"
	"flag"
	"fmt"

	"github.com/atkhx/gopangaea/internal/cli/command"
)

const (
	ChangePreset = "pc"
	CliCommand   = "change-preset"
)

func NewWithArgs(bank, preset int) (Command, error) {
	if err := Validate(bank, preset); err != nil {
		return Command{}, err
	}

	return Command{
		bank:   bank,
		preset: preset,
	}, nil
}

func New() *Command {
	return &Command{}
}

type Command struct {
	bank   int
	preset int
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %x", ChangePreset, 10*c.bank+c.preset)
}

func (c Command) GetResponseLength() int {
	return 0
}

func Validate(bank, preset int) error {
	if bank < 0 || bank > 9 {
		return errors.New("bank overflow")
	}

	if preset < 0 || preset > 9 {
		return errors.New("preset overflow")
	}
	return nil
}

func (c *Command) ParseArgs(args []string) error {
	flagset := flag.NewFlagSet("change-preset", flag.ContinueOnError)

	bank := flagset.Int("bank", 0, "index of bank [0..9]")
	preset := flagset.Int("preset", 0, "index of preset [0..9]")

	if err := flagset.Parse(args); err != nil {
		return err
	}

	if err := Validate(*bank, *preset); err != nil {
		return err
	}

	c.preset = *preset
	c.bank = *bank

	return nil
}

func (c Command) Config() command.Config {
	return command.Config{
		Command: CliCommand,
	}
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
