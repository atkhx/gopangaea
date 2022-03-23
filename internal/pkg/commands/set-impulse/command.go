package set_impulse

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

const (
	deviceCommand = "cc"
)

var ResponseSuffix = []byte{0x63, 0x63, 0x45, 0x4e, 0x44, 0x0a, 0x0d} //ccEND..

func New(name string, impulse []byte, persist bool) *Command {
	return &Command{name: name, Impulse: impulse, persist: persist}
}

type Command struct {
	name    string
	Impulse []byte
	persist bool
}

func (c Command) Validate() error {
	// todo валидация символов в имени и длины импульса
	if len(c.name) == 0 || len(c.name) > 69 {
		return errors.New("invalid name")
	}

	if len(c.Impulse) == 0 || len(c.Impulse) > 6298 {
		return errors.New("invalid impulse")
	}

	return nil
}

func (c Command) prepareImpulseData() []byte {
	result := make([]byte, len(c.Impulse)/2)

	for i := 0; i < len(c.Impulse); i += 2 {
		r, err := strconv.ParseUint(string(c.Impulse[i:i+2]), 16, 16)
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
	persist := 1
	if c.persist {
		persist = 0
	}
	return fmt.Sprintf("%s %s %d\r%x\r", deviceCommand, c.name, persist, c.prepareImpulseData())
}
