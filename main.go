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

	//setting ball to direction library
	direction.SetBall(ball)

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
<<<<<<< HEAD
			direction.MoveInDirection("N")
			time.Sleep(2000 * time.Millisecond)
			direction.MoveInDirection("S")
			time.Sleep(2000 * time.Millisecond)
=======
			direction.MoveInDirection("N", 30)
			direction.MoveInDirection("S", 30)
>>>>>>> f489b9a1187da6127d5db38aeec9edfaaca2b84a
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
