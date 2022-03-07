package exit

const (
	CliCommand     = "exit"
	CliDescription = "завершить программу"
)

type Command struct {
	stopFn func()
}

func New(stopFn func()) *Command {
	return &Command{stopFn: stopFn}
}

func (c *Command) Execute() (interface{}, error) {
	c.stopFn()
	return nil, nil
}

func (c *Command) ParseArgs(args []string) error {
	return nil
}
