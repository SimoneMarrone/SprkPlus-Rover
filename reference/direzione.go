package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

var directionPreview = "F"
// DIREZIONARE SPHERO IN AVANTI
func setDirectionForward() {
	var F := 0
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		ball.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
		})
		ball.SetBackLEDOutput(50)
		gobot.Every(1*time.Second, func() {

			if (directionPreview := "B") {
				//var B = 0
				var F = 180
				//var L = 90
				//var R = 270
				ball.Roll(0,180)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "F"
			}
			if (directionPreview := "F") {
				//var B = 180
				var F = 0
				//var L = 270
				//var R = 90
				ball.Roll(0,0)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "F"
			}
			if (directionPreview := "L") {
				//var B = 90
				var F = 270
				//var L = 180
				//var R = 0
				ball.Roll(0,270)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "F"
			}
			if (directionPreview := "R") {
				//var B = 270
				var F = 90
				//var L = 0
				//var R = 180
				ball.Roll(0,90)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "F"
			}
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
// DIREZIONARE SPHERO INDIETRO
func setDirectionBack() {
	var F := 0
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		ball.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
		})
		ball.SetBackLEDOutput(50)
		gobot.Every(1*time.Second, func() {

			if (directionPreview := "B") {
				var B = 0
				//var F = 180
				//var L = 90
				//var R = 270
				ball.Roll(0,0)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "B"
			}
			if (directionPreview := "F") {
				var B = 180
				//var F = 0
				//var L = 270
				//var R = 90
				ball.Roll(0,180)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "B"
			}
			if (directionPreview := "L") {
				var B = 90
				//var F = 270
				//var L = 180
				//var R = 0
				ball.Roll(0,90)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "B"
			}
			if (directionPreview := "R") {
				var B = 270
				//var F = 90
				//var L = 0
				//var R = 180
				ball.Roll(0,270)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "B"
			}
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
// DIREZIONARE SFERO VERSO SX
func setDirectionLeft() {
	var F := 0
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		ball.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
		})
		ball.SetBackLEDOutput(50)
		gobot.Every(1*time.Second, func() {

			if (directionPreview := "B") {
				//var B = 0
				//var F = 180
				var L = 90
				//var R = 270
				ball.Roll(0,90)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "L"
			}
			if (directionPreview := "F") {
				//var B = 180
				//var F = 0
				var L = 270
				//var R = 90
				ball.Roll(0,270)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "L"
			}
			if (directionPreview := "L") {
				//var B = 90
				//var F = 270
				var L = 180
				//var R = 0
				ball.Roll(0,180)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "L"
			}
			if (directionPreview := "R") {
				//var B = 270
				//var F = 90
				var L = 0
				//var R = 180
				ball.Roll(0,0)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "L"
			}
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
// DIREZIONARE SFERO VERSO DX
func setDirectionRight() {
	var F := 0
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	work := func() {

		ball.On("collision", func(data interface{}) {
			fmt.Printf("collision detected = %+v \n", data)
		})
		ball.SetBackLEDOutput(50)
		gobot.Every(1*time.Second, func() {

			if (directionPreview := "B") {
				//var B = 0
				//var F = 180
				//var L = 90
				var R = 270
				ball.Roll(0,270)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "R"
			}
			if (directionPreview := "F") {
				//var B = 180
				//var F = 0
				//var L = 270
				var R = 90
				ball.Roll(0,90)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "R"
			}
			if (directionPreview := "L") {
				//var B = 90
				//var F = 270
				//var L = 180
				var R = 0
				ball.Roll(0,90)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "R"
			}
			if (directionPreview := "R") {
				//var B = 270
				//var F = 90
				//var L = 0
				var R = 180
				ball.Roll(0,180)
				time.Sleep(1000 * time.Millisecond)
				directionPreview = "R"
			}
		})
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}