package get_mode

import (
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// 00000000  67 6d 0d 30 31 0a                                 |gm.01.|
const ResponseLength = 6 // gm.01.

var prefix = []byte{0x67, 0x6d, 0x0d} // gm.
var suffix = []byte{0x0a}             // .

var re = regexp.MustCompile(string(prefix) + `([\d]{2})` + string(suffix))

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
		return "phones"
	case 2:
		return "line"
	case 3:
		return "balance"
	default:
		return "unknown"
	}
}
