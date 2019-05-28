package main

import (
	"fmt"
	"os"

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
		for{
		direction.MoveInDirection("N")
		direction.MoveInDirection("S")
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
