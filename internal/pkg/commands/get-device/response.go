package get_device

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

//  61 6d 74 64 65 76 0d 30  34 0a 45 4e 44 0a        |amtdev.04.END.|
const ResponseLength = 14

var (
	prefix = []byte{0x61, 0x6d, 0x74, 0x64, 0x65, 0x76, 0x0d} // amtdev.
	suffix = []byte{0x0a, 0x45, 0x4e, 0x44, 0x0a}             // .END.
)

var re = regexp.MustCompile(string(prefix) + `(\d+)` + string(suffix))

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
		return fmt.Sprintf("pangaea cp-100 [%d]", r.Code)
	}

	return fmt.Sprintf("unknown device [%d]", r.Code)
}
