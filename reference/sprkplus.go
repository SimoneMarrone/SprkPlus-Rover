// +build example
//
// Do not build by default.

/*
 How to run
 Pass the Bluetooth address or name as the first param:

	go run examples/sprkplus.go SK-1234

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"os"
	

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	sprk := sprkplus.NewDriver(bleAdaptor)

	work := func() {
    
        sprk.SetRGB(255, 0, 0)
        sprk.Sleep()
        
    }

	robot := gobot.NewRobot("sprkBot",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{sprk},
		work,
	)

	robot.Start()
}
