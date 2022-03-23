package set_equalizer_mixer

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "eqv"
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
	// todo
	//if value < 1 || value > 200 {
	//	return errors.New("value overflow: allowed range [1..200]")
	//}
	return nil
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %d %x", deviceCommand, c.index-1, 16+c.value)
}
