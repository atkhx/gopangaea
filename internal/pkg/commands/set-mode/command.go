package set_mode

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "gm"
)

func Validate(value int) error {
	if value < 1 || value > 3 {
		return errors.New("value overflow: allowed range [1..3]")
	}
	return nil
}

func NewWithArgs(value int) (Command, error) {
	if err := Validate(value); err != nil {
		return Command{}, err
	}

	return Command{value: value}, nil
}

func New() *Command {
	return &Command{}
}

type Command struct {
	value int
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %x", deviceCommand, c.value)
}
