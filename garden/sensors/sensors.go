package sensors

import (
	"log"

	logg "../logger"
)

type Driver interface {
	Name() string
	AnalogRead(string) (int, error)
}

type Sensor struct {
	Name string
	Pin  string
	// Driver *i2c.ADS1x15Driver
	Driver Driver
}

// Init a sensor with port to board connection
// func (s *Sensor) Init(connection gpio.DigitalWriter) {
// 	s.driver = gpio.NewLedDriver(connection, s.Pin)
// }
func (s Sensor) Read() int {
	if s.Driver.Name() != "" {
		value, err := s.Driver.AnalogRead(s.Pin)
		if err != nil {
			log.Fatalln(err)
			return -1
		}

		logg.SaveData(s.Name, value)
		log.Println("Sensor:", s.Name, " Pin:", s.Pin, " Value:", value)
		return value
	}

	value := 0 // random value
	log.Println("Sensor:", s.Name, " Pin:", s.Pin, " Value:", value)
	return value
}

// Init Sensors
func InitSensors(driver Driver) []Sensor {
	sensors := []Sensor{
		{
			Name: "soilTemp",
			Pin:  "0",
		},
		{
			Name: "soilMoist",
			Pin:  "1",
		},
		{
			Name: "airTemp",
			Pin:  "2",
		},
		{
			Name: "airLight",
			Pin:  "3",
		},
	}

	for i := range sensors {
		sensors[i].Driver = driver
	}
	return sensors
}

// ReadSensors reads the values of all sensors
func ReadSensors(sensors []Sensor) {

	for _, sensor := range sensors {
		sensor.Read()
	}
}
