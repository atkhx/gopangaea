package command

type Command interface {
	GetCommand() string
	GetResponseLength() int
	ParseResponse(b []byte) (interface{}, error)
	ParseArgs(args []string) error
}

type Config struct {
	Command   string
	Arguments map[string]string
}
