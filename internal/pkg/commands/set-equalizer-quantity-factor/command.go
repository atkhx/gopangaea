package set_equalizer_quantity_factor

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "eqq"
)

func New(index, value int) Command {
	return Command{value: value, index: index}
}

type Command struct {
	index int
	value int
}

func (c Command) Validate() error {
	if c.index < 1 || c.index > 5 {
		return errors.New("index overflow: allowed range [1..5]")
	}
	if c.value < 1 || c.value > 200 {
		return errors.New("value overflow: allowed range [1..200]")
	}
	return nil
}

func (c Command) GetCommand() string {
	value := c.value - 101
	if value < 0 {
		value = value + 256
	}
	return fmt.Sprintf("%s %d %x", deviceCommand, c.index-1, value)
}
