package set_settings

import "fmt"

const (
	deviceCommand = "gs"
)

type Command struct {
	Settings []byte
}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s 1\r%x\r", deviceCommand, c.Settings)
}

func (c Command) GetResponseLength() int {
	return len(successResponse)
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
