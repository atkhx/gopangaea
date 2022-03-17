package get_settings

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"

	"github.com/atkhx/gopangaea/internal/pkg/valreader"
)

const ResponseLength = 90

type Response struct {
	Bytes []byte
}

func ParseResponse(data []byte) (Response, error) {
	if len(data) != ResponseLength {
		return Response{}, errors.New(fmt.Sprintf("invalid data length: %d", len(data)))
	}
	return Response{Bytes: data[3:]}, nil
}

func (r Response) String() string {
	return hex.Dump(r.Bytes)
}

func (r Response) Settings() Settings {
	//fmt.Println(hex.Dump(r.Bytes))

	lpv := valreader.ReadIntFromChar(string(r.Bytes[72:74]), "LowPassFilter.Value")
	lpv = int(math.Pow(195.0-float64(lpv), 2)*(19000.0/math.Pow(195.0, 2.0)) + 1000.0)

	hpv := valreader.ReadIntFromChar(string(r.Bytes[74:76]), "HighPassFilter.Value")
	hpv = (hpv*980)/255 + 20

	prh := valreader.ReadIntFromChar(string(r.Bytes[38:40]), "PreAmp.Equalizer.high") - 192
	if prh < 0 {
		prh = 256 + prh
	}

	prm := valreader.ReadIntFromChar(string(r.Bytes[36:38]), "PreAmp.Equalizer.mid") - 192
	if prm < 0 {
		prm = 256 + prm
	}

	prl := valreader.ReadIntFromChar(string(r.Bytes[34:36]), "PreAmp.Equalizer.low") - 192
	if prl < 0 {
		prl = 256 + prl
	}

	return Settings{
		NoiseGate: NoiseGate{
			Active: valreader.ReadIntFromChar(string(r.Bytes[40:42]), "NoiseGate.Active") > 0,
			Thresh: valreader.ReadIntFromChar(string(r.Bytes[42:44]), "NoiseGate.Thresh"),
			Decay:  valreader.ReadIntFromChar(string(r.Bytes[44:46]), "NoiseGate.Decay"),
		},
		Compressor: Compressor{
			Active:  valreader.ReadIntFromChar(string(r.Bytes[46:48]), "Compressor.Active") > 0,
			Sustain: valreader.ReadIntFromChar(string(r.Bytes[48:50]), "Compressor.Sustain"),
			Volume:  valreader.ReadIntFromChar(string(r.Bytes[50:52]), "Compressor.Volume"),
		},
		PreAmp: PreAmp{
			Active:    valreader.ReadIntFromChar(string(r.Bytes[30:32]), "PreAmp.Active") > 0,
			Volume:    valreader.ReadIntFromChar(string(r.Bytes[32:34]), "PreAmp.Volume"),
			Equalizer: [3]int{prl, prm, prh},
		},
		EarlyReverb: EarlyReverb{
			Active: valreader.ReadIntFromChar(string(r.Bytes[20:22]), "EarlyReverb.Active") > 0,
			Volume: valreader.ReadIntFromChar(string(r.Bytes[10:12]), "EarlyReverb.Volume"),
			Type:   valreader.ReadIntFromChar(string(r.Bytes[12:14]), "EarlyReverb.Type"),
		},
		PowerAmp: PowerApm{
			Active: valreader.ReadIntFromChar(string(r.Bytes[22:24]), "PowerAmp.Active") > 0,
			Volume: valreader.ReadIntFromChar(string(r.Bytes[24:26]), "PowerAmp.Volume"),
			Slave:  valreader.ReadIntFromChar(string(r.Bytes[26:28]), "PowerAmp.Slave"),
			Index:  valreader.ReadIntFromChar(string(r.Bytes[28:30]), "PowerAmp.Index"),
		},
		Presence: Presence{
			Active: valreader.ReadIntFromChar(string(r.Bytes[80:82]), "Presence.Active") > 0,
			Value:  valreader.ReadIntFromChar(string(r.Bytes[82:84]), "Presence.Value"),
		},
		Impulse: Impulse{
			Active: valreader.ReadIntFromChar(string(r.Bytes[16:18]), "Impulse") > 0,
		},
		MasterVolume: valreader.ReadIntFromChar(string(r.Bytes[14:16]), "MasterVolume"),
		Equalizer: Equalizer{
			Active:   valreader.ReadIntFromChar(string(r.Bytes[18:20]), "Equalizer.Active") > 0,
			Position: valreader.ReadIntFromChar(string(r.Bytes[84:86]), "Equalizer.Position"),
			Mixer: [5]int{
				mixer(valreader.ReadIntFromChar(string(r.Bytes[0:2]), "Equalizer.Mixer1")),
				mixer(valreader.ReadIntFromChar(string(r.Bytes[2:4]), "Equalizer.Mixer2")),
				mixer(valreader.ReadIntFromChar(string(r.Bytes[4:6]), "Equalizer.Mixer3")),
				mixer(valreader.ReadIntFromChar(string(r.Bytes[6:8]), "Equalizer.Mixer4")),
				mixer(valreader.ReadIntFromChar(string(r.Bytes[8:10]), "Equalizer.Mixer5")),
			},
			Frequencies: [5]int{
				frequence(120, 1, valreader.ReadIntFromChar(string(r.Bytes[52:54]), "Equalizer.FR1")),
				frequence(360, 1, valreader.ReadIntFromChar(string(r.Bytes[54:56]), "Equalizer.FR2")),
				frequence(800, 2, valreader.ReadIntFromChar(string(r.Bytes[56:58]), "Equalizer.FR3")),
				frequence(2000, 10, valreader.ReadIntFromChar(string(r.Bytes[58:60]), "Equalizer.FR4")),
				frequence(6000, 50, valreader.ReadIntFromChar(string(r.Bytes[60:62]), "Equalizer.FR5")),
			},
			QualityFactor: [5]int{
				quantityFactory(valreader.ReadIntFromChar(string(r.Bytes[62:64]), "Equalizer.QF1")),
				quantityFactory(valreader.ReadIntFromChar(string(r.Bytes[64:66]), "Equalizer.QF2")),
				quantityFactory(valreader.ReadIntFromChar(string(r.Bytes[66:68]), "Equalizer.QF2")),
				quantityFactory(valreader.ReadIntFromChar(string(r.Bytes[68:70]), "Equalizer.QF2")),
				quantityFactory(valreader.ReadIntFromChar(string(r.Bytes[70:72]), "Equalizer.QF2")),
			},
		},
		LowPassFilter: LowPassFilter{
			Active: valreader.ReadIntFromChar(string(r.Bytes[78:80]), "LowPassFilter.Active") > 0,
			Value:  lpv,
		},
		HighPassFilter: HighPassFilter{
			Active: valreader.ReadIntFromChar(string(r.Bytes[76:78]), "HighPassFilter.Active") > 0,
			Value:  hpv,
		},
	}
}

func quantityFactory(i int) int {
	res := 101 + i
	if res > 256 {
		res -= 256
	}
	return res
}

func frequence(def, koef, val int) int {
	if val > 100 {
		val -= 256
	}
	return def + koef*(val)
}

func mixer(val int) int {
	val = val - 16
	if val < 0 {
		val += 1
	}
	return val
}
