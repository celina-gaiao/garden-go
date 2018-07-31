package main

import (
	"fmt"
	"time"

	act "./actuators"
	"./api"
	sens "./sensors"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/chip"
)

func main() {

	// var wg sync.WaitGroup

	chipAdaptor := chip.NewProAdaptor()

	// default bus/address
	ads1115 := i2c.NewADS1115Driver(chipAdaptor)

	// optional bus/address
	// ads1115 := i2c.NewADS1015Driver(chipAdaptor,
	// 	i2c.WithBus(1),
	// 	i2c.WithAddress(0x48)) -> see CHIP for address

	sensors := sens.InitSensors(ads1115)

	solenoids := act.InitSolenoids()
	for i := range solenoids {
		solenoids[i].Init(chipAdaptor)
	}

	solenoidDrivers := make([]*gpio.LedDriver, len(solenoids))
	for i, solenoid := range solenoids {
		solenoidDrivers[i] = solenoid.Driver
	}

	work := func() {
		gobot.Every(1*time.Minute, func() {
			act.WaterPlants(solenoids)
			sens.ReadSensors(sensors)
		})
	}

	robot := gobot.NewRobot("GardenBot",
		[]gobot.Connection{chipAdaptor},
		solenoidDrivers,
		[]gobot.Device{ads1115},
		work,
	)

	go robot.Start()

	fmt.Println("start api")

	api.Start()

}
