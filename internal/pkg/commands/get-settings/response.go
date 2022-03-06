package get_settings

import (
	"encoding/hex"
	"errors"
	"fmt"
)

const ResponseLength = 90

type Response struct {
	Bytes []byte
}

func ParseResponse(data []byte) (Response, error) {
	if len(data) != ResponseLength {
		return Response{}, errors.New(fmt.Sprintf("invalid data length: %d", len(data)))
	}
	return Response{Bytes: data[:]}, nil
}

func (r Response) String() string {
	return hex.Dump(r.Bytes)
}
