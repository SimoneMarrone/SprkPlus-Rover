// +build example
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
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		ball.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
		})
		ball.SetBackLEDOutput(50)
		gobot.Every(1*time.Second, func() {
			ball.Roll(50, 0)
			time.Sleep(2600 * time.Millisecond)
			ball.Roll(0, 180)
			time.Sleep(1000 * time.Millisecond)
			ball.Roll(50, 180)
			time.Sleep(2600 * time.Millisecond)
			ball.Roll(0, 0)
			time.Sleep(1000 * time.Millisecond)
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
