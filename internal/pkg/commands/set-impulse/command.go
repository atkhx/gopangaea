package set_impulse

import (
	"errors"
	"fmt"
)

const (
	deviceCommand = "cc"
)

var ResponseSuffix = []byte{0x63, 0x63, 0x45, 0x4e, 0x44, 0x0a, 0x0d} //ccEND..

func New(name string, persist bool, impulse Impulse) *Command {
	return &Command{name: name, impulse: impulse, persist: persist}
}

type Command struct {
	name    string
	persist bool

	impulse Impulse
}

func (c Command) Validate() error {
	// todo валидация символов в имени и длины импульса
	if len(c.name) == 0 || len(c.name) > 69 {
		return errors.New("invalid name")
	}

	return c.impulse.IsValid()
}

func (c Command) GetCommand() string {
	// 0 - перезатирает пресет
	// - передаем имя
	// - передем заголови
	// 1 - временное хранение
	// - s вместо имени
	// - отрезаем заголовки
	if c.persist {
		return fmt.Sprintf("%s %s %d\r%x\r", deviceCommand, c.name, 0, c.impulse.Source())
	} else {
		trimmed, err := c.impulse.Trimmed()
		if err != nil {
			panic("run without validation")
		}
		return fmt.Sprintf("%s %s %d\r%x\r", deviceCommand, "s", 1, trimmed)
	}
}
