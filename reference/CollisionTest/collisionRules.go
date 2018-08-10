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
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

const ms float64 = 0.4347826

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)

	work := func() {

		//Start time in millisecond
		ball.Roll(52, 0)
		start := time.Now()

		ball.On("collision", func(data interface{}) {

			//Get time on collision in second
			elapsed := time.Since(start)

			fmt.Println("Secondi: ", elapsed)
			elapsed.Seconds()
			//Get meter travel
			mRide := elapsed.Seconds() * ms
			fmt.Printf("raw data: %f \n", elapsed.Seconds())
			if elapsed.Seconds() <= 1.38 {
				fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.187))
			} else if 1.38 <= elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
				fmt.Printf("elapsed 30cm= %f \n", mRide-(elapsed.Seconds()*0.14))
			} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.70 {
				fmt.Printf("elapsed 50cm= %f \n", mRide-(elapsed.Seconds()*0.11))
			} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
				fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
				fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			} else {
				fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			}

			ball.Stop()

		})
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
