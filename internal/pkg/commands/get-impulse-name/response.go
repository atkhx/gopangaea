package get_impulse_name

import (
	"regexp"

	"github.com/pkg/errors"
)

// 00000000  72 6e 0d 6d 61 5f 5f 5f  5f 5f 5f 5f 5f 5f 5f 5f  |rn.ma___________|
// 00000010  5f 5f 5f 31 39 36 30 41  56 5f 5f 5f 5f 5f 5f 5f  |___1960AV_______|
// 00000020  5f 5f 5f 34 78 31 32 5f  56 69 6e 74 61 67 65 5f  |___4x12_Vintage_|
// 00000030  33 30 5f 43 6f 6e 64 5f  5f 5f 5f 43 61 70 5f 30  |30_Cond____Cap_0|
// 00000040  5f 5f 2e 77 61 76 0a                              |__.wav.|

const ResponseLength = 71

var (
	prefix = []byte{0x72, 0x6e, 0x0d} // rn.
	suffix = []byte{0x0a}             // .
)

var re = regexp.MustCompile(string(prefix) + `(\*|[\S]{1,63}\.wav)` + string(suffix))

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
