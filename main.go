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
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	Info.Println(fmt.Sprintf("Hey, we are running on %s...\n", runtime.GOOS))

	check(checkUser())

	keyboards := findKeyboards()

	Info.Println(fmt.Sprintf("Keyboards in the system: %v\n", extractDeviceNames(keyboards)))

}