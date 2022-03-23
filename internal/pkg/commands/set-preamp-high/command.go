package set_preamp_high

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "prh"
)

func New(value int) Command {
	return Command{value: value}
}

type Command struct {
	value int
}

func (c Command) Validate() error {
	if c.value < 0 || c.value > 127 {
		return errors.New("value overflow: allowed range [0..127]")
	}
	return nil
}
func (c Command) GetCommand() string {
	v := 192 + c.value
	if v > 256 {
		v = v - 256
	}
	return fmt.Sprintf("%s %x", deviceCommand, byte(v))
}
