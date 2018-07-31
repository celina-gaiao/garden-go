package sensors

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

func TestCanReadOneSensor(t *testing.T) {
	// test Read()

	sensor := Sensor{
		Name:   "test",
		Pin:    "40",
		Driver: DriverMock{},
	}

	value := sensor.Read()
	if value == 0 {
		t.Errorf("Read error: %s", string(value))
	}
}

func TestCanReadSeveralSensors(t *testing.T) {
	// test ReadSensors

	sensors := []Sensor{
		{
			Name:   "test1",
			Pin:    "41",
			Driver: DriverMock{},
		},
		{
			Name:   "test2",
			Pin:    "42",
			Driver: DriverMock{},
		},
		{
			Name:   "test3",
			Pin:    "43",
			Driver: DriverMock{},
		},
	}

	values := ReadSensors(sensors)
	fmt.Println(values)

	if values == nil {
		t.Errorf("No values read")
	}

	for _, v := range values {
		if v != 88 {
			t.Errorf("Unespected value")
		}
	}

}
