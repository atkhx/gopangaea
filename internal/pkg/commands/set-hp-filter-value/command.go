package set_hp_filter_value

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "hv"
)

func Validate(value int) error {
	if value < 20 || value > 1000 {
		return errors.New("value overflow: allowed range [20..1000]")
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
	// Y = (X*980)/255 + 20
	// X = (255 * (Y - 20))/980
	val := int(255*(c.value-20)) / 980
	return fmt.Sprintf("%s %x", deviceCommand, val)
}
