package get_impulse_name

import (
	"regexp"

	"github.com/pkg/errors"
)

const ResponseLength = 69

var (
	suffix = []byte{0x0a} // .
)

var re = regexp.MustCompile(`(\*|[\S]{1,64}\.wav)` + string(suffix))

type Response struct {
	Name string
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 2 {
		return Response{}, errors.New("parse error")
	}
	return Response{Name: string(matches[1])}, nil
}

func (r Response) String() string {
	return r.Name
}
