package set_reverb_type

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "et"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 0 || c.value > 2 {
		return errors.New("value overflow: allowed range [0..2]")
	}
	return nil
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %d", deviceCommand, c.value)
}
