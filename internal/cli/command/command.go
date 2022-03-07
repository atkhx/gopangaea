package command

type CliCommand interface {
	Execute() (interface{}, error)
	ParseArgs(args []string) error
}
