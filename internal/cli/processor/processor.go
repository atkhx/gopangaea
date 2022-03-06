package processor

import (
	"github.com/atkhx/gopangaea/internal/cli/command"
	"github.com/atkhx/gopangaea/internal/pkg/device"
)

type Device interface {
	ExecCommand(command device.Command) ([]byte, error)
}

func New(device Device) *processor {
	return &processor{device: device}
}

type processor struct {
	device Device
}

func (p *processor) Exec(in command.Command) (interface{}, error) {
	r, err := p.device.ExecCommand(in)
	if err != nil {
		return nil, err
	}
	return in.ParseResponse(r)
}
