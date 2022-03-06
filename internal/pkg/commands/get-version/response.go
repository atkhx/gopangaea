package get_version

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const ResponseLength = 15 // amtver.6.08.04.

var prefix = []byte{0x61, 0x6d, 0x74, 0x76, 0x65, 0x72, 0x0d} // amtver.
var suffix = []byte{0x0a}                                     // .

// todo сделать валидацию длинны в регулярке
var re = regexp.MustCompile(string(prefix) + `(\d+)\.(\d+)\.(\d+)` + string(suffix))

type Response struct {
	Major int
	Minor int
	Patch int
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 4 {
		return Response{}, errors.New("parse error")
	}

	major, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse major version failed")
	}

	minor, err := strconv.Atoi(string(matches[2]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse minor version failed")
	}

	patch, err := strconv.Atoi(string(matches[3]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse patch version failed")
	}

	return Response{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func (r Response) String() string {
	return fmt.Sprintf("%d.%d.%d", r.Major, r.Minor, r.Patch)
}
