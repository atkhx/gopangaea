package set_mode

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "gm"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 1 || c.value > 3 {
		return errors.New("value overflow: allowed range [1..3]")
	}
	return nil
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %x", deviceCommand, c.value)
}
