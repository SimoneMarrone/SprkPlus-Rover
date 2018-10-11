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

func directionNE() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("\n turn NE")
		ball.Roll(0, 45)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 45)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionNO() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("\n turn NO")
		ball.Roll(0, 315)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 315)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionSE() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("\n turn SE")
		ball.Roll(0, 135)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 135)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionSO() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("\n turn SO")
		ball.Roll(0, 225)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 225)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionO() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Print("\n turn left")
		ball.Roll(0, 270)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 270)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionE() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("\n turn right")
		ball.Roll(0, 90)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 90)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionN() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("\n turn forward")
		ball.Roll(0, 0)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 0)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
func directionS() {

	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)
	work := func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print("\n turn back")
		ball.Roll(0, 180)
		time.Sleep(500 * time.Millisecond)
		ball.Roll(50, 180)
	}
	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
