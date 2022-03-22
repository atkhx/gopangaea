package set_impulse

import (
	"fmt"
	"log"
	"strconv"
)

const (
	deviceCommand = "cc"
)

var ResponseSuffix = []byte{0x63, 0x63, 0x45, 0x4e, 0x44, 0x0a, 0x0d} //ccEND..

type Command struct {
	Name    string
	Impulse []byte
}

func New() *Command {
	return &Command{}
}

func convertResponse(data []byte) []byte {
	result := make([]byte, len(data)/2)

	for i := 0; i < len(data); i += 2 {
		r, err := strconv.ParseUint(string(data[i:i+2]), 16, 16)
		if err != nil {
			log.Fatalln(err)
		}
		result[i/2] = byte(r)
	}

	return result
}

func (c Command) GetCommand() string {
	// 0 - перезатирает пресет
	// 1 - временное хранение
	return fmt.Sprintf("%s %s 1\r%x\r", deviceCommand, c.Name, convertResponse(c.Impulse))
}
