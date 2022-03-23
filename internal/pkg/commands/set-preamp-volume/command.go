package set_preamp_volume

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "prv"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 0 || c.value > 31 {
		return errors.New("value overflow: allowed range [0..31]")
	}
	return nil
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %x", deviceCommand, c.value)
}
