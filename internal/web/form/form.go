package form

import (
	get_settings "github.com/atkhx/gopangaea/internal/pkg/commands/get-settings"
)

func FromDto(dto get_settings.Settings) Form {
	return Form{
		NoiseGate: NoiseGate{
			Active: dto.NoiseGate.Active,
			Thresh: dto.NoiseGate.Thresh,
			Decay:  dto.NoiseGate.Decay,
		},
		Compressor: Compressor{
			Active:  dto.Compressor.Active,
			Sustain: dto.Compressor.Sustain,
			Volume:  dto.Compressor.Volume,
		},
		PreAmp: PreAmp{
			Active:    dto.PreAmp.Active,
			Volume:    dto.PreAmp.Volume,
			Equalizer: dto.PreAmp.Equalizer,
		},
		EarlyReverb: EarlyReverb{
			Active: dto.EarlyReverb.Active,
			Volume: dto.EarlyReverb.Volume,
			Type:   dto.EarlyReverb.Type,
		},
		PowerAmp: PowerApm{
			Active: dto.PowerAmp.Active,
			Volume: dto.PowerAmp.Volume,
			Slave:  dto.PowerAmp.Slave,
			Index:  dto.PowerAmp.Index,
		},
		AmpTypes: GetAmpTypes(dto.PowerAmp.Index),
		Presence: Presence{
			Active: dto.Presence.Active,
			Value:  dto.Presence.Value,
		},
		Impulse: Impulse{
			Active: dto.Impulse.Active,
		},
		MasterVolume: dto.MasterVolume,
		HighPassFilter: HighPassFilter{
			Active: dto.HighPassFilter.Active,
			Value:  dto.HighPassFilter.Value,
		},
		LowPassFilter: LowPassFilter{
			Active: dto.LowPassFilter.Active,
			Value:  dto.LowPassFilter.Value,
		},
		Equalizer: Equalizer{
			Active:        dto.Equalizer.Active,
			Position:      dto.Equalizer.Position,
			Mixer:         dto.Equalizer.Mixer,
			Frequencies:   dto.Equalizer.Frequencies,
			QualityFactor: dto.Equalizer.QualityFactor,
		},
	}
}

type Form struct {
	NoiseGate    NoiseGate
	Compressor   Compressor
	PreAmp       PreAmp
	EarlyReverb  EarlyReverb
	PowerAmp     PowerApm
	Presence     Presence
	Impulse      Impulse
	Mode         int
	MasterVolume int

	AmpTypes []ampType

	Equalizer      Equalizer
	LowPassFilter  LowPassFilter
	HighPassFilter HighPassFilter
}

type Equalizer struct {
	Active        bool
	Position      int    // pre / post
	Mixer         [5]int // EQ-Микшер
	Frequencies   [5]int // EQ-Частоты
	QualityFactor [5]int // EQ-Добротность
}

type EarlyReverb struct {
	Active bool
	Volume int
	Type   int
}

type Impulse struct {
	Active bool
}

type PowerApm struct {
	Active bool
	Volume int
	Slave  int
	Index  int
	Type   int
}

type PreAmp struct {
	Active    bool
	Volume    int
	Equalizer [3]int
}

type Presence struct {
	Active bool
	Value  int
}

type NoiseGate struct {
	Active bool
	Thresh int
	Decay  int
}

type Compressor struct {
	Active  bool
	Sustain int
	Volume  int
}

type LowPassFilter struct {
	Active bool
	Value  int
}

type HighPassFilter struct {
	Active bool
	Value  int
}
