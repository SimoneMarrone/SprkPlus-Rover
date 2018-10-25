package main

import (
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

	//setting ball to direction library
	direction.SetBall(ball)

	//keeping library
	prolog.AssertBusy("L")

	//map init
	maps.InitMap()
	maps.PrintMap()

	//main function for sprk
	work := func() {

		//taking direction
		direction.DirectionNE()
	}

	//New adapter
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()

}
