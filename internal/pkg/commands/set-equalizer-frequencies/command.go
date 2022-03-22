package set_equalizer_frequencies

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "eqf"
)

func Validate(index, value int) error {
	if index < 1 || index > 5 {
		return errors.New("index overflow: allowed range [1..5]")
	}
	// todo разные частоты
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
	var def, koef int
	switch c.index - 1 {
	case 0:
		def, koef = 120, 1
	case 1:
		def, koef = 360, 1
	case 2:
		def, koef = 800, 2
	case 3:
		def, koef = 2000, 10
	case 4:
		def, koef = 6000, 50
	}

	val := 256 + (c.value-def)/koef
	return fmt.Sprintf("%s %d %x", deviceCommand, c.index-1, val)
}
