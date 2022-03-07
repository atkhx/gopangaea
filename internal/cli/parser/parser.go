package parser

import (
	"errors"
	"strings"

	"github.com/atkhx/gopangaea/internal/cli/command"
)

var (
	ErrCommandIsEmpty   = errors.New("empty command")
	ErrCommandIsUnknown = errors.New("unknown command")
)

func New(knownCommands map[string]command.CliCommand) *parser {
	return &parser{knownCommands: knownCommands}
}

type parser struct {
	knownCommands map[string]command.CliCommand
}

func (p *parser) Parse(input string) (command.CliCommand, error) {
	args := strings.Split(input, " ")
	for k, v := range args {
		args[k] = strings.TrimSpace(v)
	}

	if len(args) == 0 || args[0] == "" {
		return nil, ErrCommandIsEmpty
	}

	cmd, isKnownCommand := p.knownCommands[args[0]]
	if !isKnownCommand {
		return nil, ErrCommandIsUnknown
	}

	if err := cmd.ParseArgs(args[1:]); err != nil {
		return nil, err
	}

	return cmd, nil
}
