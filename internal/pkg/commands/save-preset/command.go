package save_reset

const (
	deviceCommand = "sp"
)

func New() *Command {
	return &Command{}
}

type Command struct {
}

func (c Command) GetCommand() string {
	return deviceCommand
}
