package set_poweramp_type

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "at"
)

func Validate(value int) error {
	if value < 0 || value > 15 {
		return errors.New("value overflow: allowed range [0..14]")
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
	fmt.Println(fmt.Sprintf("%s %d", deviceCommand, c.value))
	return fmt.Sprintf("%s %d", deviceCommand, c.value)
}

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
