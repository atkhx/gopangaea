package set_hp_filter_value

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "hv"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 20 || c.value > 1000 {
		return errors.New("value overflow: allowed range [20..1000]")
	}
	return nil
}

func (c Command) GetCommand() string {
	// Y = (X*980)/255 + 20
	// X = (255 * (Y - 20))/980
	val := int(255*(c.value-20)) / 980
	return fmt.Sprintf("%s %x", deviceCommand, val)
}
