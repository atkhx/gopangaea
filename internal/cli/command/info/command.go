package info

import (
	"bytes"
	"fmt"

	"github.com/jpoirier/gousb/usb"
)

const (
	CliCommand     = "info"
	CliDescription = "показать информацию о подключении"
)

type Command struct {
	dev *usb.Device
}

func New(dev *usb.Device) *Command {
	return &Command{dev: dev}
}

func (c *Command) Execute() (interface{}, error) {
	s, err := c.dev.GetStringDescriptor(2)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(nil)

	desc := c.dev.Descriptor

	buffer.WriteString(fmt.Sprintf("device descriptor:       %s\n", s))
	buffer.WriteString(fmt.Sprintf("device dev.desc.Product: %s\n", desc.Product))
	buffer.WriteString(fmt.Sprintf("device dev.desc.Vendor:  %s\n", desc.Vendor))

	return buffer.String(), nil
}

func (c *Command) ParseArgs(args []string) error {
	return nil
}
