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

func (c Command) GetResponseLength() int {
	return len(successResponse)
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
