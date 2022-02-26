package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var files = []string{
	"preset_0_0",
	"preset_1_0",
	"preset_2_1",
	"preset_3_2",
}

func main() {
	for _, file := range files {
		panFileName := fmt.Sprintf("/Users/aatikhonov/go/src/github.com/atkhx/gopangaea/cmd/test/%s.pan", file)
		dumpFileName := fmt.Sprintf("/Users/aatikhonov/go/src/github.com/atkhx/gopangaea/cmd/test/%s.dump", file)

		b, err := ioutil.ReadFile(panFileName)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(hex.Dump(b))

		if err := ioutil.WriteFile(dumpFileName, []byte(hex.Dump(b)), os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}
