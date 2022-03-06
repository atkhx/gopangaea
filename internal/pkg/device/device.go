package device

import (
	"github.com/atkhx/gopangaea/internal/pkg/commands/change-preset"
	get_bank "github.com/atkhx/gopangaea/internal/pkg/commands/get-bank"
	get_device "github.com/atkhx/gopangaea/internal/pkg/commands/get-device"
	get_impulce_name "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-name"
	get_impulse_names "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-names"
	get_mode "github.com/atkhx/gopangaea/internal/pkg/commands/get-mode"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	get_version "github.com/atkhx/gopangaea/internal/pkg/commands/get-version"
)

type Command interface {
	GetCommand() string
	GetResponseLength() int
}

type CommandWriter interface {
	Write(command string) error
}

type ResponseReader interface {
	Read() ([]byte, error)
	ReadWithSkipTails(command string, length int) ([]byte, error)
}

type device struct {
	writer CommandWriter
	reader ResponseReader
}

func New(writer CommandWriter, reader ResponseReader) *device {
	return &device{writer: writer, reader: reader}
}

func (d *device) ExecCommand(command Command) ([]byte, error) {
	return d.execCommand(command.GetCommand(), command.GetResponseLength())
}

func (d *device) execCommand(command string, responseLength int) ([]byte, error) {
	if err := d.writer.Write(command); err != nil {
		return nil, err
	}

	return d.reader.ReadWithSkipTails(command, responseLength)
}

func (d *device) GetDevice() (get_device.Response, error) {
	b, err := d.ExecCommand(get_device.Command{})
	if err != nil {
		return get_device.Response{}, err
	}
	return get_device.ParseResponse(b)
}

func (d *device) GetVersion() (get_version.Response, error) {
	b, err := d.ExecCommand(get_version.Command{})
	if err != nil {
		return get_version.Response{}, err
	}
	return get_version.ParseResponse(b)
}

func (d *device) GetBank() (get_bank.Response, error) {
	b, err := d.ExecCommand(get_bank.Command{})
	if err != nil {
		return get_bank.Response{}, err
	}
	return get_bank.ParseResponse(b)
}

func (d *device) GetImpulseName() (get_impulce_name.Response, error) {
	b, err := d.ExecCommand(get_impulce_name.Command{})
	if err != nil {
		return get_impulce_name.Response{}, err
	}
	return get_impulce_name.ParseResponse(b)
}

func (d *device) GetMode() (get_mode.Response, error) {
	b, err := d.ExecCommand(get_mode.Command{})
	if err != nil {
		return get_mode.Response{}, err
	}
	return get_mode.ParseResponse(b)
}

func (d *device) GetSettings() (get_settings.Response, error) {
	b, err := d.ExecCommand(get_settings.Command{})
	if err != nil {
		return get_settings.Response{}, err
	}
	return get_settings.ParseResponse(b)
}

func (d *device) GetImpulseNames() (get_impulse_names.Response, error) {
	b, err := d.ExecCommand(get_impulse_names.Command{})
	if err != nil {
		return get_impulse_names.Response{}, err
	}
	return get_impulse_names.ParseResponse(b)
}

func (d *device) ChangePreset(bank, preset int) (change_preset.Response, error) {
	command, err := change_preset.NewWithArgs(bank, preset)
	if err != nil {
		return change_preset.Response{}, err
	}

	b, err := d.ExecCommand(command)
	if err != nil {
		return change_preset.Response{}, err
	}

	return change_preset.ParseResponse(b)
}
