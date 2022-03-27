package get_impulse

import (
	"encoding/hex"
	"regexp"

	"github.com/atkhx/gopangaea/internal/pkg/library/impulse"
	"github.com/pkg/errors"
)

const ResponseLength = 69 + 6298 + 1

var (
	suffix = []byte{0x0a} // .
)

var re = regexp.MustCompile(`(\*|[\S]{1,64}\.wav)` + string(suffix) + `(.*?)` + string(suffix))

type Response struct {
	Name    string
	Impulse *impulse.Impulse
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 3 {
		return Response{}, errors.New("parse error")
	}

	originBin := make([]byte, hex.DecodedLen(len(matches[2])))
	n, err := hex.Decode(originBin, matches[2])
	if err != nil {
		return Response{}, err
	}
	if n == 0 {
		return Response{}, errors.New("empty oiginBin")
	}

	originImpulse := impulse.New(originBin[:n])

	if err := originImpulse.IsValid(); err != nil {
		return Response{}, errors.Wrap(err, "invalid impulse")
	}

	response := Response{
		Name:    string(matches[1]),
		Impulse: originImpulse,
	}

	return response, nil
}

func (r Response) String() string {
	return r.Name
}
