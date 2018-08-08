// build example
//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth address or name as the first param:

	go run examples/bb8-collision.go BB-1234

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	dir := 0

	work := func() {
		ball.Stop()

		ball.On("collision", func(data interface{}) {
			fmt.Println("Collision Detected!")
			ball.SetRGB(uint8(gobot.Rand(255)),
				uint8(gobot.Rand(255)),
				uint8(gobot.Rand(255)),
			)

			ball.Roll(10, uint16(-dir))

			//change direction
			dir = gobot.Rand(360)
			fmt.Println(ball.Name())

		})

		gobot.Every(3*time.Second, func() {
			ball.Roll(30, uint16(dir))
		})

	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
