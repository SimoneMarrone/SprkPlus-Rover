package direction

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/platforms/sphero/ollie"
)

var ball *ollie.Driver

var wait time.Duration = 1000
var wait_2 time.Duration = 500

func SetBall(main_ball *ollie.Driver) {
	ball = main_ball
}

func DirectionNE() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn NE")
	ball.Roll(0, 45)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 45)

	//on.collision{
	//SET nuova posizione = passare a map
	//direzione presa
	//distanza percorsa

	//SET ostacolo su mappa
	//posizione?

	//}

	//IF NO COLLISION
	//SET nuova posizione = passare a map
	//direzione presa
	//distanza percorsa

}

func DirectionNO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn NO")
	ball.Roll(0, 315)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 315)

}
func DirectionSE() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn SE")
	ball.Roll(0, 135)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 135)

}
func DirectionSO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn SO")
	ball.Roll(0, 225)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 225)

}
func DirectionO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn left")
	ball.Roll(0, 270)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 270)

}
func DirectionE() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn right")
	ball.Roll(0, 90)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 90)

}

func DirectionN() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn forward")
	ball.Roll(0, 0)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 0)
}
func DirectionS() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn back")
	ball.Roll(0, 180)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(50, 180)

}
