package set_poweramp_type

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "at"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 0 || c.value > 15 {
		return errors.New("value overflow: allowed range [0..14]")
	}
	return nil
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %d", deviceCommand, c.value)
}
