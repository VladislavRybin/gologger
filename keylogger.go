package main

import (
	"os"
	"fmt"
	"encoding/binary"
	"bytes"
)

func NewKeyLogger(dev *InputDevice) *KeyLogger {
	return &KeyLogger{
		dev: dev,
	}
}

func (t *KeyLogger) Read() (chan InputEvent, error) {
	ret := make(chan InputEvent, 512)

	fd, err := os.Open(fmt.Sprintf(DEVICE_FILE, t.dev.Id))
	if err != nil {
		close(ret)
		return ret, fmt.Errorf("Error opening device file:", err)
	}

	go func() {

		tmp := make([]byte, eventsize)
		event := InputEvent{}
		for {

			n, err := fd.Read(tmp)
			if err != nil {
				panic(err)
				close(ret)
				break
			}
			if n <= 0 {
				continue
			}

			if err := binary.Read(bytes.NewBuffer(tmp), binary.LittleEndian, &event); err != nil {
				panic(err)
			}

			ret <- event

		}
	}()
	return ret, nil
}

func (ev *InputEvent) KeyString() string {
	return keyCodeMap[ev.Code]
}