package set_settings

import "fmt"

const (
	deviceCommand = "gs"
)

// "gsEND.."
var ResponseSuffix = []byte{0x67, 0x73, 0x45, 0x4e, 0x44, 0x0a, 0x0d}

type Command struct {
	Settings []byte
}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s 1\r%x\r", deviceCommand, c.Settings)
}
