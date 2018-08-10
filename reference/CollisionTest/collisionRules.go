/**
	Calculation centimeters after collision

	INFO:
	100cm in 2300ms = 0.4347826 m/s
**/

package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

const ms float64 = 0.434782

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		//Start time in millisecond
		start := time.Now()

		ball.On("collision", func(data interface{}) {

			//Get time on collision in second
			elapsed := time.Since(start)

			//Get meter travel
			mRide := elapsed.Seconds() * ms
			fmt.Printf("elapsed = %f \n", mRide)
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
