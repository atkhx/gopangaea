package change_preset

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "pc"
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
	return fmt.Sprintf("%s %x", deviceCommand, 10*c.bank+c.preset)
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
