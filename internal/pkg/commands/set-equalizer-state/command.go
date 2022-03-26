package set_equalizer_state

import (
	"fmt"
)

const (
	deviceCommand = "eqo"
)

func New(value bool) Command {
	return Command{value: value}
}

type Command struct {
	value bool
}

func (c Command) GetCommand() string {
	if c.value {
		return fmt.Sprintf("%s 1", deviceCommand)
	}
	return fmt.Sprintf("%s 0", deviceCommand)
}
