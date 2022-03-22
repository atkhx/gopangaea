package set_hp_filter_state

import (
	"fmt"
)

const (
	deviceCommand = "ho"
)

func NewWithArgs(value bool) Command {
	return Command{value: value}
}

func New() *Command {
	return &Command{}
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
