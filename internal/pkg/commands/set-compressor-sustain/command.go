package set_compressor_sustain

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "cs"
)

func Validate(value int) error {
	if value < 0 || value > 31 {
		return errors.New("value overflow: allowed range [0..31]")
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

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
