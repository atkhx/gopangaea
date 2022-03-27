package reset_preset

import (
	"fmt"
)

const (
	deviceCommand = "esc"
)

func New() Command {
	return Command{}
}

type Command struct {
}

func (c Command) GetCommand() string {
	return fmt.Sprintf("%s", deviceCommand)
}
