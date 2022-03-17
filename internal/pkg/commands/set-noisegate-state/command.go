package set_noisegate_state

import (
	"fmt"
)

const (
	deviceCommand = "go"
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

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
