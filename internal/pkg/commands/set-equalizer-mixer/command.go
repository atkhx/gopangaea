package set_equalizer_mixer

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "eqv"
)

func Validate(index, value int) error {
	if index < 1 || index > 5 {
		return errors.New("index overflow: allowed range [1..5]")
	}
	// todo
	//if value < 1 || value > 200 {
	//	return errors.New("value overflow: allowed range [1..200]")
	//}
	return nil
}

func NewWithArgs(index, value int) (Command, error) {
	if err := Validate(index, value); err != nil {
		return Command{}, err
	}

	return Command{value: value, index: index}, nil
}

func New() *Command {
	return &Command{}
}

type Command struct {
	index int
	value int
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s %d %x", deviceCommand, c.index-1, 16+c.value)
}

func (c Command) GetResponseLength() int {
	return 0
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
