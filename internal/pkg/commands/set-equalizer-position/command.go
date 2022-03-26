package set_equalizer_position

import (
	"fmt"
)

const (
	deviceCommand = "eqp"
)

func New(pre bool) Command {
	return Command{pre: pre}
}

type Command struct {
	pre bool
}

func (c Command) GetCommand() string {
	if c.pre {
		return fmt.Sprintf("%s 1", deviceCommand)
	}
	return fmt.Sprintf("%s 0", deviceCommand)
}
