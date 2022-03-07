package usage

import (
	"bytes"
	"sort"
)

const (
	CliCommand     = "help"
	CliDescription = "показать доступные команды"
)

type Command struct {
	knownCommands map[string]string
}

func New(knownCommands map[string]string) *Command {
	return &Command{knownCommands: knownCommands}
}

func (c *Command) Execute() (interface{}, error) {
	var commands []string
	for command := range c.knownCommands {
		commands = append(commands, command)
	}

	sort.Strings(commands)

	buffer := bytes.NewBuffer(nil)

	buffer.WriteString("Usage:\n\n")
	buffer.WriteString("\tcommand [arguments]\n\n")

	buffer.WriteString("Available commands:\n\n")
	for _, command := range commands {
		descriptionSpace := "\t"
		switch {
		case len(command) < 8:
			descriptionSpace = "\t\t\t"
		case len(command) < 16:
			descriptionSpace = "\t\t"
		}
		buffer.WriteString("\t" + command + descriptionSpace + c.knownCommands[command] + "\n")
	}

	return buffer.String(), nil
}

func (c *Command) ParseArgs(args []string) error {
	return nil
}
