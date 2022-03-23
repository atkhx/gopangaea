package change_preset

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "pc"
)

func New(bank, preset int) Command {
	return Command{
		bank:   bank,
		preset: preset,
	}
}

type Command struct {
	bank   int
	preset int
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %x", deviceCommand, 10*c.bank+c.preset)
}

func (c Command) Validate() error {
	if c.bank < 0 || c.bank > 9 {
		return errors.New("bank overflow")
	}

	if c.preset < 0 || c.preset > 9 {
		return errors.New("preset overflow")
	}
	return nil
}
