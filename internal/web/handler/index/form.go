package index

import (
	"fmt"

	get_impulse_names "github.com/atkhx/gopangaea/internal/pkg/commands/get-impulse-names"
	form "github.com/atkhx/gopangaea/internal/pkg/library/amp"
	get_settings "github.com/atkhx/gopangaea/internal/pkg/library/settings"
)

func DeviceImpulsesFromDto(response get_impulse_names.Response) (result DeviceImpulses) {
	for _, name := range response.Names {
		result = append(result, DeviceImpulse{
			Name:     name,
			Selected: false,
		})
	}
	return
}

type Form struct {
	Settings get_settings.Settings
	AmpTypes []AmpType
}

type DeviceImpulses []DeviceImpulse

type DeviceImpulse struct {
	Name     string
	Selected bool
}

type AmpTypes []AmpType

type AmpType struct {
	Name     string
	Selected bool
}

func GetAmpTypes(version string, selected int) (res []AmpType) {
	for i, amplifier := range form.GetAmplifiers(version) {
		res = append(res, AmpType{
			Name:     amplifier,
			Selected: selected == i,
		})
	}
	return
}

type PresetSelection struct {
	Banks   []PresetSelectionRow
	Presets []PresetSelectionRow
}

type PresetSelectionRow struct {
	Label    string
	Selected bool
}

func GetPresetSelection(currentBank, currentPreset int) (res PresetSelection) {
	for bank := 0; bank < 10; bank++ {
		res.Banks = append(res.Banks, PresetSelectionRow{
			Label:    fmt.Sprintf("%d", bank),
			Selected: currentBank == bank,
		})
	}

	for preset := 0; preset < 10; preset++ {
		res.Presets = append(res.Presets, PresetSelectionRow{
			Label:    fmt.Sprintf("%d", preset),
			Selected: currentPreset == preset,
		})
	}

	return
}
