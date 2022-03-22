package get_impulse

const (
	deviceCommand = "cc"
)

type Command struct {
}

func New() *Command {
	return &Command{}
}

func (c Command) GetCommand() string {
	return deviceCommand
}

func (c Command) GetResponseLength() int {
	return ResponseLength
}

func (c Command) ParseResponse(b []byte) (interface{}, error) {
	return ParseResponse(b)
}
