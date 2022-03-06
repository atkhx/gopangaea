package get_settings

type Settings struct {
	Equalizer      Equalizer
	EarlyReverb    EarlyReverb
	MasterVolume   byte
	Impulse        Impulse
	PowerApm       PowerApm
	PreAmp         PreAmp
	Presence       Presence
	NoiseGate      NoiseGate
	Compressor     Compressor
	LowPassFilter  LowPassFilter
	HighPassFilter HighPassFilter
}

type Equalizer struct {
	Active        bool
	Position      byte    // pre / post
	Mixer         [5]byte // EQ-Микшер
	Frequencies   [5]byte // EQ-Частоты
	QualityFactor [5]byte // EQ-Добротность
}

type EarlyReverb struct {
	Active bool
	Volume byte
	Type   byte
}

type Impulse struct {
	Active bool
}

type PowerApm struct {
	Active bool
	Volume byte
	Slave  byte
	Index  byte
}

type PreAmp struct {
	Active    bool
	Volume    byte
	Equalizer [3]byte
}

type Presence struct {
	Active bool
	Value  byte
}

type NoiseGate struct {
	Active bool
	Thresh byte
	Decay  byte
}

type Compressor struct {
	Active  bool
	Sustain byte
	Volume  byte
}

type LowPassFilter struct {
	Active bool
	Value  byte
}

type HighPassFilter struct {
	Active bool
	Value  byte
}
