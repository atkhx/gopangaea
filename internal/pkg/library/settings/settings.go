package settings

type Settings struct {
	NoiseGate   NoiseGate
	Compressor  Compressor
	PreAmp      PreAmp
	EarlyReverb EarlyReverb
	PowerAmp    PowerApm
	Presence    Presence

	ImpulseState bool
	MasterVolume int

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

type PowerApm struct {
	Active bool
	Volume int
	Slave  int
	Index  int
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
