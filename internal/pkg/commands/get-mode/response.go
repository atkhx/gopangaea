package get_mode

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const ResponseLength = 3 // 01.

var suffix = []byte{0x0a} // .

var re = regexp.MustCompile(`([\d]{2})` + string(suffix))

type Response struct {
	Mode int
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 2 {
		return Response{}, errors.New("parse error")
	}

	mode, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse mode version failed")
	}

	return Response{Mode: mode}, nil
}

func (r Response) String() string {
	switch r.Mode {
	case 1:
		return "Phones"
	case 2:
		return "Line"
	case 3:
		return "Balance"
	default:
		return "Unknown"
	}
}
