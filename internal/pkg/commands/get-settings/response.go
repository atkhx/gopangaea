package get_settings

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

const ResponseLength = 87

type Response struct {
	Bytes []byte
}

func ParseResponse(data []byte) (Response, error) {
	if len(data) != ResponseLength {
		return Response{}, errors.New(fmt.Sprintf("invalid data length: %d", len(data)))
	}
	return Response{Bytes: convertResponse(data[:86])}, nil
}

func (r Response) String() string {
	return hex.Dump(r.Bytes)
}

func convertResponse(data []byte) []byte {
	result := make([]byte, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		r, err := strconv.ParseUint(string(data[i:i+2]), 16, 16)
		if err != nil {
			log.Fatalln(err)
		}
		result[i/2] = byte(r)
	}

	return result
}

func (r Response) Settings() Settings {
	lpv := int(math.Pow(195.0-float64(int(r.Bytes[36])), 2)*(19000.0/math.Pow(195.0, 2.0)) + 1000.0)
	hpv := (int(r.Bytes[37])*980)/255 + 20

	prh := int(r.Bytes[19]) - 192
	if prh < 0 {
		prh = 256 + prh
	}

	prm := int(r.Bytes[18]) - 192
	if prm < 0 {
		prm = 256 + prm
	}

	prl := int(r.Bytes[17]) - 192
	if prl < 0 {
		prl = 256 + prl
	}

	return Settings{
		NoiseGate: NoiseGate{
			Active: r.Bytes[20] > 0,
			Thresh: int(r.Bytes[21]),
			Decay:  int(r.Bytes[22]),
		},
		Compressor: Compressor{
			Active:  r.Bytes[23] > 0,
			Sustain: int(r.Bytes[24]),
			Volume:  int(r.Bytes[25]),
		},
		PreAmp: PreAmp{
			Active:    r.Bytes[15] > 0,
			Volume:    int(r.Bytes[16]),
			Equalizer: [3]int{prl, prm, prh},
		},
		EarlyReverb: EarlyReverb{
			Active: r.Bytes[10] > 0,
			Volume: int(r.Bytes[5]),
			Type:   int(r.Bytes[6]),
		},
		PowerAmp: PowerApm{
			Active: r.Bytes[11] > 0,
			Volume: int(r.Bytes[12]),
			Slave:  int(r.Bytes[13]),
			Index:  int(r.Bytes[14]),
		},
		Presence: Presence{
			Active: r.Bytes[40] > 0,
			Value:  int(r.Bytes[41]),
		},
		Impulse: Impulse{
			Active: r.Bytes[8] > 0,
		},
		MasterVolume: int(r.Bytes[7]),
		Equalizer: Equalizer{
			Active:   r.Bytes[9] > 0,
			Position: int(r.Bytes[42]),
			Mixer: [5]int{
				mixer(int(r.Bytes[0])),
				mixer(int(r.Bytes[1])),
				mixer(int(r.Bytes[2])),
				mixer(int(r.Bytes[3])),
				mixer(int(r.Bytes[4])),
			},
			Frequencies: [5]int{
				frequence(120, 1, int(r.Bytes[26])),
				frequence(360, 1, int(r.Bytes[27])),
				frequence(800, 2, int(r.Bytes[28])),
				frequence(2000, 10, int(r.Bytes[29])),
				frequence(6000, 50, int(r.Bytes[30])),
			},
			QualityFactor: [5]int{
				quantityFactory(int(r.Bytes[31])),
				quantityFactory(int(r.Bytes[32])),
				quantityFactory(int(r.Bytes[33])),
				quantityFactory(int(r.Bytes[34])),
				quantityFactory(int(r.Bytes[35])),
			},
		},
		LowPassFilter: LowPassFilter{
			Active: r.Bytes[39] > 0,
			Value:  lpv,
		},
		HighPassFilter: HighPassFilter{
			Active: r.Bytes[38] > 0,
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
