package get_bank

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

const ResponseLength = 5 // 0300.

var suffix = []byte{0x0a} // .

var re = regexp.MustCompile(`0(\d)0(\d)` + string(suffix))

type Response struct {
	Bank   int
	Preset int
}

func ParseResponse(data []byte) (Response, error) {
	matches := re.FindSubmatch(data)
	if len(matches) != 3 {
		return Response{}, errors.New("parse error")
	}

	bank, err := strconv.Atoi(string(matches[1]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse bank index failed")
	}

	preset, err := strconv.Atoi(string(matches[2]))
	if err != nil {
		return Response{}, errors.Wrap(err, "parse preset index failed")
	}

	return Response{Bank: bank, Preset: preset}, nil
}

func (r Response) String() string {
	return fmt.Sprintf("bank %d preset %d", r.Bank, r.Preset)
}
