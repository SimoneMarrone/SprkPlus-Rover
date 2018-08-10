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
	"os"
    "time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/sprkplus"
)

func main() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
    
    var wait time.Duration 
    var speed uint8
    
	work := func() {
		
        
        for ok := true; ok; ok = true {
            
            wait = 1000
            speed = 45
            
            
            
            ball.Roll(speed,0)
            ball.SetRGB(0, 255, 0)
            time.Sleep(wait * time.Millisecond)
            ball.Roll(speed,90)
            ball.SetRGB(255, 0, 0)
            time.Sleep(wait * time.Millisecond)
            ball.Roll(speed,180)
            ball.SetRGB(0, 255, 0)
            time.Sleep(wait * time.Millisecond)
            ball.Roll(speed,270)
            ball.SetRGB(255, 0, 0)
            time.Sleep(wait * time.Millisecond)
            ball.SetRGB(255, 0, 0)
            
        }

		
		
  
	}

	robot := gobot.NewRobot("sprkplus",
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{ball},
		work,
	)

	robot.Start()
}
