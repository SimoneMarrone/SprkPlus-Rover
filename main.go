package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/ble"
	"gobot.io/x/gobot/platforms/sphero/ollie"

	"./prolog"

	"./reference/Direction"
	"./reference/Maps"
)

func main() {

	var isCollision = false

	//get parameters from console (SK-9885)
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := ollie.NewDriver(bleAdaptor)

	// Default configuration
	ball.EnableStopOnDisconnect()
	ball.SetBackLEDOutput(1)
	ball.SetStabilization(true)
	ball.SetRotationRate(1)
	ball.On("collision", func(data interface{}) {
		isCollision = true

		/* Colore RGB - rosso */
		ball.SetRGB(255, 0, 0)

		/* Tempo di collisione */
		elapsed := time.Since(direction.Start)

		/* Cm percorsi */
		mRide := elapsed.Seconds() * direction.Ms
		fmt.Printf("Tempo di collisione %f \n", (mRide / 10)

		for i := 0; i < int(mRide / 10); i++ {
			//setPosition(direction)
		}
		//maps.SetObstacle()
	})

	//setting ball to direction library
	direction.SetBall(ball)

	//map init
	Maps.InitMap()
	Maps.PrintMap()

	work := func() {
		for {
			prolog.SetDirOfMap()
			direction.Start = time.Now()

			speed, direct = prolog.MakeMove()
			if !isCollision {
				for i := 0; i < speed; i++ {
					direction.SetPosition(direct)
				}
			}

			prolog.Reset()
		}
	}

	if len(os.Args) > 2 {
		switch os.Args[2] {
		case "no-blt":
			for {
				work()
			}
		case "interactive":
			direction.Wait = 500
			work := func() {
				for {
					buf := bufio.NewReader(os.Stdin)
					fmt.Print("> ")
					sentence, _ := buf.ReadString('\n')
					sentence = strings.Replace(sentence, "\n", "", -1)
					fmt.Println(sentence)
					direction.MoveInDirection(string(sentence), 30)
				}
			}
			//New adapter
			robot := gobot.NewRobot("sprkplus",
				[]gobot.Connection{bleAdaptor},
				[]gobot.Device{ball},
				work)

			robot.Start()

			break
		}
	} else { //normal execution

		//New adapter
		robot := gobot.NewRobot("sprkplus",
			[]gobot.Connection{bleAdaptor},
			[]gobot.Device{ball},
			work,
		)

		robot.Start()
	}

}
