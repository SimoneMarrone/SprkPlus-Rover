package direction

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/platforms/sphero/ollie"
)

const ms float64 = 0.4347826
const wait time.Duration = 500
const speed uint8 = 50

/* DRIVER */
var ball *ollie.Driver

func SetBall(mainBall *ollie.Driver) {
	ball = mainBall
}

func getGradeDirection(grade string) uint16 {
	switch grade {
	case "NE":
		return 45
	case "NW":
		return 315
	case "SE":
		return 135
	case "SW":
		return 225
	case "W":
		return 270
	case "E":
		return 90
	case "N":
		return 0
	case "S":
		return 180
	default:
		return 0
	}
}

func moveInDirection(direction string) {

	/* Movimento del robot */
	time.Sleep(wait * time.Millisecond)
	ball.Roll(0, getGradeDirection(direction))
	time.Sleep(wait * time.Millisecond)
	ball.Roll(speed, getGradeDirection(direction))

	start := time.Now()
	ball.On("collision", func(data interface{}) {

		/* Tempo di collisione */
		elapsed := time.Since(start)
		elapsed.Seconds()
		/* Metri percorsi */
		mRide := elapsed.Seconds() * ms
		fmt.Printf("Percorso %f \n", mRide-(elapsed.Seconds()*0.092))
	})
}
