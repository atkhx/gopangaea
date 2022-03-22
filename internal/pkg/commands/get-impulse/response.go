package get_impulse

import (
	"regexp"

	"github.com/pkg/errors"
)

const ResponseLength = 69 + 6298 + 1

var (
	suffix = []byte{0x0a} // .
)

var re = regexp.MustCompile(`(\*|[\S]{1,64}\.wav)` + string(suffix) + `(.*?)` + string(suffix))

type Response struct {
	Name    string
	Impulse []byte
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 3 {
		return Response{}, errors.New("parse error")
	}

	response := Response{
		Name:    string(matches[1]),
		Impulse: matches[2],
	}

	return response, nil
}

func (r Response) String() string {
	return r.Name
}
