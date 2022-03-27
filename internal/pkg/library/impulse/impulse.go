package impulse

import (
	"bytes"
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/youpy/go-wav"
)

const numSamplesCP100 = 984

func New(origin []byte) *Impulse {
	return &Impulse{origin: origin}
}

type Impulse struct {
	origin []byte
}

func (i *Impulse) IsValid() error {
	reader := wav.NewReader(bytes.NewReader(i.origin))
	format, err := reader.Format()
	if err != nil {
		return err
	}

	if format.NumChannels != 1 {
		return errors.Errorf("invalid NumChannels: %d", format.NumChannels)
	}

	return nil
}

func (i *Impulse) Trimmed() ([]byte, error) {
	reader := wav.NewReader(bytes.NewReader(i.origin))

	format, err := reader.Format()
	if err != nil {
		return nil, err
	}

	samples, err := reader.ReadSamples()
	if err != nil && err != io.EOF {
		return nil, err
	}

	result := bytes.NewBuffer([]byte{})
	wrt := wav.NewWriter(result, uint32(numSamplesCP100), format.NumChannels, format.SampleRate, format.BitsPerSample)

	result.Reset()

	if err := wrt.WriteSamples(samples[:numSamplesCP100]); err != nil {
		log.Fatalln(err)
	}

	return result.Bytes(), nil
}

func (i *Impulse) Source() []byte {
	result := make([]byte, len(i.origin))
	copy(result, i.origin)
	return result
}
