package set_lp_filter_value

import (
	"errors"
	"fmt"
	"math"
)

const (
	deviceCommand = "lv"
)

func Validate(value int) error {
	if value < 1000 || value > 20000 {
		return errors.New("value overflow: allowed range [1k..20k]")
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
	// Y = (195-X)^2 * (19000/195^2) + 1000
	// X = 195 - sqrt((Y - 1000) / (19000/195^2))
	val := int(195 - math.Sqrt((float64(c.value-1000))/(19000/math.Pow(195, 2))))

	return fmt.Sprintf("%s %x", deviceCommand, val)
}
