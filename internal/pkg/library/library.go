package library

import (
	get_impulse "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
	"github.com/pkg/errors"
)

type Device interface {
	ChangePreset(bank, preset int) (bool, error)
	GetImpulse() (get_impulse.Response, error)
	GetSettings() (get_settings.Response, error)
}

func New() *library {
	return &library{}
}

type library struct {
	presets []Preset
}

func (l *library) GetPreset(index int) (Preset, error) {
	if index < 0 || index > len(l.presets) {
		return Preset{}, errors.New("preset index overflow")
	}

	return l.presets[index], nil
}

func (l *library) LoadFromDevice(device Device) error {
	var presets []Preset
	for bank := 0; bank < 10; bank++ {
		for p := 0; p < 10; p++ {
			_, err := device.ChangePreset(bank, p)
			if err != nil {
				return errors.Wrap(err, "change preset failed")
			}

			i, err := device.GetImpulse()
			if err != nil {
				return errors.Wrap(err, "get impulse failed")
			}

			settings, err := device.GetSettings()
			if err != nil {
				return errors.Wrap(err, "get settings failed")
			}

			presets = append(presets, Preset{
				Name:     i.Name,
				Settings: settings.Settings(),
				Impulse:  i.Impulse,
			})
		}
	}

	l.presets = presets
	return nil
}
