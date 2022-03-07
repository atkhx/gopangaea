package set_lp_filter_value

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "lv"
)

func Validate(value int) error {
	// todo провалидировать min/max
	if value < 0 || value > 255 {
		return errors.New("value overflow: allowed range [0..255]")
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
