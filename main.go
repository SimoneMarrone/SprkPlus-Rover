package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/ollie"

	"./prolog"

	"./reference/Direction"
	"./reference/Maps"
)

func main() {

	//get parameters from console (SK-9885)
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)

	// Default configuration
	ball.EnableStopOnDisconnect()
	ball.SetBackLEDOutput(1)
	ball.SetStabilization(true)
	ball.SetRotationRate(1)
	ball.On("collision", func(data interface{}) {

		/* Colore RGB - rosso */
		ball.SetRGB(255, 0, 0)

		/* Tempo di collisione */
		elapsed := time.Since(start)

		/* Metri percorsi */
		//mRide := (elapsed.Seconds() * ms) - elapsed.Seconds()
		mRide := elapsed.Seconds() * ms
		fmt.Printf("Tempo di collisione %f \n", mRide)

		for i := 0; i < interval; i++ {
			setPosition(direction)
		}
		//maps.SetObstacle()
	})

	//setting ball to direction library
	direction.SetBall(ball)

	//keeping library
	prolog.AssertBusy("L")

	//map init
	maps.InitMap()
	maps.PrintMap()

	//main function for sprk
	work := func() {

		//prolog.AssertBusy("w")
		//fmt.Printf("%t",prolog.CheckDirection("w"))

		prolog.SetDirOfMap()
		fmt.Printf("%s", prolog.FreeDir())

		//fmt.Printf("%s",prolog.FreeDir())
		//taking direction
		for {
			direction.MoveInDirection("N", 30)
			direction.MoveInDirection("S", 30)
		}
	}

	//New adapter
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()

}
