package valreader

import (
	"log"
	"strconv"
)

func ReadIntFromChar(value string, s ...interface{}) int {
	r, err := strconv.ParseUint(value, 16, 16)
	if err != nil {
		log.Fatalln(err)
	}
	if len(s) > 0 {
		//s = append(s, "value", r, "source", value)
		//fmt.Println(s...)
	}
	return int(r)
}
