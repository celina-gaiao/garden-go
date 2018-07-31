package actuators

import (
	"log"
	"strconv"
	"strings"
	"time"

	logg "../logger"

	"gobot.io/x/gobot/drivers/gpio"
)

type Scheduler struct {
	Start    string // start time defined as 10:00 or 14:00
	Duration int    // in minutes
}

type Solenoid struct {
	Name      string
	Pin       string
	Driver    *gpio.LedDriver
	Schedules []Scheduler
}

// Open a solenoid for watering
func (s Solenoid) OpenFor(duration time.Duration) {
	s.Open()
	<-time.After(duration)
	s.Close()
}

// Init a solenoid with port to board connection
func (s *Solenoid) Init(connection gpio.DigitalWriter) {
	s.Driver = gpio.NewLedDriver(connection, s.Pin)
	// TODO this needs to be done on work???
	s.Driver.Off()
}

// Open a solenoid for watering
func (s Solenoid) Open() {
	log.Printf("watering %s\n", s.Name)
	s.Driver.On()
	logg.SaveData(s.Name, 1)
}

// Close a solenoid to stop watering
func (s Solenoid) Close() {
	log.Printf("stop watering %s\n", s.Name)
	s.Driver.Off()
	logg.SaveData(s.Name, 0)
}

func InitSolenoids() []Solenoid {

	solenoids := []Solenoid{
		{
			Name: "pot",
			Pin:  "CSID7",
			Schedules: []Scheduler{
				{Start: "7:23", Duration: 1},
				{Start: "18:05", Duration: 1},
			},
		},
		{
			Name: "bed",
			Pin:  "CSID6",
			Schedules: []Scheduler{
				{Start: "7:08", Duration: 10},
				{Start: "18:01", Duration: 5},
			},
		},
		{
			Name: "tube",
			Pin:  "CSID5",
			Schedules: []Scheduler{
				{Start: "7:12", Duration: 15},
				{Start: "18:10", Duration: 10},
			},
		},
		{
			Name: "empty",
			Pin:  "CSID4",
			//Schedules: []Scheduler{
			//	{Start: "10:00", Duration: 1},
			//},
		},
	}
	return solenoids
}

// WaterPlants cheks if it is time to water and make it water
func WaterPlants(solenoids []Solenoid) {

	for _, solenoid := range solenoids {
		for _, schedule := range solenoid.Schedules {

			parts := strings.Split(schedule.Start, ":")
			hour, _ := strconv.Atoi(parts[0])
			minute, _ := strconv.Atoi(parts[1])

			currentHour, currentMinute, _ := time.Now().Clock()

			if hour == currentHour && minute == currentMinute {
				// go func(s Solenoid, d int) {
				// 	s.OpenFor(time.Duration(d) * time.Minute)
				// }(solenoid, schedule.Duration)
				go solenoid.OpenFor(time.Duration(schedule.Duration) * time.Minute)
			}
		}
	}
}
