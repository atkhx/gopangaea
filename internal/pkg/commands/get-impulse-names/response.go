package get_impulse_names

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const (
	PresetCountMax            = 100
	PresetImpulseMaxLength    = 74
	PresetImpulseSuffixLength = 8                                                  // .01.END.
	PresetMinLength           = 1 + PresetImpulseSuffixLength                      // *.01.END.
	PresetMaxLength           = PresetImpulseMaxLength + PresetImpulseSuffixLength // name.01.END.

	ResponseLength    = PresetMaxLength * PresetCountMax
	MinResponseLength = PresetMinLength * PresetCountMax
)

var (
	separator = []byte{0x0a, 0x45, 0x4e, 0x44, 0x0a} // .END.
)

var re = regexp.MustCompile(`(\*|[\S]{1,64}\.wav)` + string([]byte{0x0a}) + `\d{2}`)

type Response struct {
	Names []string
}

func ParseResponse(data []byte) (Response, error) {
	if len(data) < MinResponseLength {
		return Response{}, errors.New("invalid data for parse")
	}

	var response = Response{}
	content := bytes.Split(data, separator)
	for i, v := range content {
		if matches := re.FindSubmatch(v); len(matches) > 1 {
			response.Names = append(response.Names, string(matches[1]))
		} else if len(v) > 0 {
			return Response{}, fmt.Errorf("parse impulse [%d] in `%s` failed", i, string(v))
		}
	}
	return response, nil
}

func (r Response) String() string {
	return strings.Join(r.Names, "\n")
}
