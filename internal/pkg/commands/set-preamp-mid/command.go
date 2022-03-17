package set_preamp_mid

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "prm"
)

func Validate(value int) error {
	if value < 0 || value > 127 {
		return errors.New("value overflow: allowed range [0..127]")
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
	v := 192 + c.value
	if v > 256 {
		v = v - 256
	}
	return fmt.Sprintf("%s %x", deviceCommand, byte(v))
}

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
