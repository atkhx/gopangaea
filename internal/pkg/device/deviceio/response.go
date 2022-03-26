package deviceio

import (
	"encoding/hex"

	"github.com/pkg/errors"
)

// "END."
var responseWithEND = []byte{0x45, 0x4e, 0x44, 0x0a}

// "00."
var responseWith00 = []byte{0x30, 0x30, 0x0a}

// "01."
var responseWith01 = []byte{0x30, 0x31, 0x0a}

func NewResponse(suffix ...[]byte) *ResponseBool {
	return &ResponseBool{
		possible: append([][]byte{
			responseWithEND,
			responseWith00,
			responseWith01,
		}, suffix...),
	}
}

type ResponseBool struct {
	success  bool
	possible [][]byte
}

func (r *ResponseBool) GetLength() (result int) {
	for _, suffix := range r.possible {
		if l := len(suffix); result == 0 || l < result {
			result = l
		}
	}
	return
}

func (r *ResponseBool) Parse(actual []byte) error {
	if len(actual) == 0 {
		r.success = true
		return nil
	}

	for _, suffix := range r.possible {
		if r.success = string(actual) == string(suffix); r.success {
			return nil
		}
	}

	return errors.Errorf("invalid response: %s", hex.Dump(actual))
}

func (r ResponseBool) Success() bool {
	return r.success
}

func (r ResponseBool) String() string {
	if r.success {
		return "success"
	}
	return "failed"
}
