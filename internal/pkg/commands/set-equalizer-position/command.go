package set_equalizer_position

import (
	"fmt"
)

const (
	deviceCommand = "eqp"
)

func NewWithArgs(pre bool) Command {
	return Command{pre: pre}
}

func New() *Command {
	return &Command{}
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

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
