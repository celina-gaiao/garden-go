package logger

import (
	"fmt"
	"testing"
)

type DriverMock struct {
}

func (d DriverMock) Name() string {
	return "mock"
}

func (d DriverMock) AnalogRead(pin string) (value int, err error) {
	value = 88
	err = nil

	return
}

func TestCanSaveDataUnit(t *testing.T) {
	// test SaveData()

	var data []Data

	SaveData("temperature", 20)

	data = GetData()

	fmt.Println(data)

}
