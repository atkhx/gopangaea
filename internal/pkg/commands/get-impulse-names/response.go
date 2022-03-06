package get_impulse_names

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//00000000  1b 5b 32 4a 1b 5b 48  72 6e 73  0d 2a 0a 30 31 0a  |rns.*.01.|
//00000000  1b 5b 32 4a 1b 5b 48  72 6e 73  0d 2a 0a 30 31 0a  |.[2J.[Hrns.*.01.|
//00000000  27 91 50 74 27 91 72 114
//00000000   . 91 50 74  . 91 72   r  n  s
//00000000   .  [ 50 74  .  [ 72

// fe______________65_TWN__________2x12_Jensen_C12KCond____Cap_2__.wav.01.END - 74
//00000000  72 6e 73 0d 46 65 6e 64  65 72 5f 36 35 5f 54 77  |rns.Fender_65_Tw|
//00000010  69 6e 5f 32 78 31 32 5f  4a 65 6e 73 65 6e 5f 43  |in_2x12_Jensen_C|
//00000020  31 32 4b 43 6f 6e 64 5f  43 61 70 5f 30 5f 2e 77  |12KCond_Cap_0_.w|
//00000030  61 76 0a 30 31 0a 45 4e  44 0a 66 65 5f 5f 5f 5f  |av.01.END.fe____|
//00000040  5f 5f 5f 5f 5f 5f 5f 5f  5f 5f 36 35 5f 54 57 4e  |__________65_TWN|
//00000050  5f 5f 5f 5f 5f 5f 5f 5f  5f 5f 32 78 31 32 5f 4a  |__________2x12_J|
//00000060  65 6e 73 65 6e 5f 43 31  32 4b 43 6f 6e 64 5f 5f  |ensen_C12KCond__|
//00000070  5f 5f 43 61 70 5f 32 5f  5f 2e 77 61 76 0a 30 31  |__Cap_2__.wav.01|
//00000080  0a 45 4e 44 0a 2a 0a 30  31 0a 45 4e 44 0a 2a 0a  |.END.*.01.END.*.|

const (
	PresetCountMax            = 100
	PresetImpulseMaxLength    = 74
	ResponsePrefixLength      = 4                                                  // rns.
	PresetImpulseSuffixLength = 8                                                  // .01.END.
	PresetMinLength           = 1 + PresetImpulseSuffixLength                      // *.01.END.
	PresetMaxLength           = PresetImpulseMaxLength + PresetImpulseSuffixLength // name.01.END.

	ResponseLength    = ResponsePrefixLength + PresetMaxLength*PresetCountMax
	MinResponseLength = ResponsePrefixLength + PresetMinLength*PresetCountMax
)

var (
	prefix    = []byte{0x72, 0x6e, 0x73, 0x0d} // rns.
	separator = []byte{0x0a, 0x45, 0x4e, 0x44, 0x0a}
)

var re = regexp.MustCompile(`(\*|[\S]{1,63}\.wav)` + string([]byte{0x0a}) + `\d{2}`)

type Response struct {
	Names []string
}

func ParseResponse(data []byte) (Response, error) {
	if len(data) < MinResponseLength || !bytes.HasPrefix(data, prefix) {
		fmt.Println(hex.Dump(data))
		return Response{}, errors.New("invalid data for parse")
	}

	var response = Response{}
	content := bytes.Split(data[len(prefix):], separator)
	for i, v := range content {
		if matches := re.FindSubmatch(v); len(matches) > 1 {
			response.Names = append(response.Names, string(matches[1]))
		} else if len(v) > 0 {
			return Response{}, fmt.Errorf("parse impulse [%d] in `%s` failed", i, string(v))
		}
	}
	return response, nil
}

func (r Response) String() string {
	return strings.Join(r.Names, "\n")
}
