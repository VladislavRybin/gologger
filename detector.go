package main

import (
	"os/user"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func checkUser() error {
	u, err := user.Current()
	if err != nil {
		return err
	}

	Info.Println(fmt.Sprintf("Current user name is %s...\n", u.Username))

	if u.Uid != "0" {
		Error.Println("Cannot read device files. Are you running as root?\n")
		return fmt.Errorf("Check privileges!\n")
	}
	return nil
}

func extractDeviceNames(arr []int) []string {
	var names []string
	for i:=0; i < len(arr); i++ {
		buff, err := ioutil.ReadFile(fmt.Sprintf(UEVENT_FILE, i))
		check(err)

		names = append(names, strings.Split(strings.Split(string(buff), "\n")[1], "=")[1])
	}
	return names
}

func findKeyboards() []int {
	var keyboard_indices []int

	devices := 0
	keyboards := 0
	for i := 0; i< MAX_FILES; i++ {

		var buff []byte
		var err error

		if (fileExists(fmt.Sprintf(UEVENT_FILE, i))) {
			buff, err = ioutil.ReadFile(fmt.Sprintf(UEVENT_FILE, i))
			check(err)
		}

		if isKeyboard(string(buff)) {
			keyboard_indices = append(keyboard_indices, i)
			keyboards++
		}

		devices++
	}
	Info.Println(fmt.Sprintf("Input devices in the system: %d\n", devices))

	return keyboard_indices
}

func isKeyboard(dvc string) bool {
	if strings.Contains(dvc, "EV="); strings.Contains(dvc, "KEY=") {
		arr := strings.Split(dvc, "\n")

		for j := 0; j < len(arr); j++ {
			if strings.Contains(arr[j], "KEY=") {

				key := strings.Split(arr[j], "=")[1]
				words := strings.Split(key, " ")

				return strings.HasSuffix(words[len(words)-1], "fffffffe")
			}
		}
	}
	return false
}