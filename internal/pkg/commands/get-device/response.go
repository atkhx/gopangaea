package get_device

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

// 04.END.
const ResponseLength = 7

var (
	suffix = []byte{0x0a, 0x45, 0x4e, 0x44, 0x0a} // .END.
)

var re = regexp.MustCompile(`(\d+)` + string(suffix))

type Response struct {
	Code int
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 2 {
		return Response{}, errors.New("parse error")
	}

	code, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse code failed")
	}

	return Response{Code: code}, nil
}

func (r Response) String() string {
	if r.Code == 4 {
		return "Pangaea CP-100"
	}

	return fmt.Sprintf("unknown device [%d]", r.Code)
}
