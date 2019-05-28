package direction

import (
	"fmt"
	"time"

	"../Maps"
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

const ms float64 = 0.4347826
const wait time.Duration = 3000
const interval int = 5

/* DRIVER */
var ball *ollie.Driver

func SetBall(mainBall *ollie.Driver) {
	ball = mainBall
}

func getGradeDirection(direction string) uint16 {
	switch direction {
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

func setPosition(direction string) {
	switch direction {
	case "NE":
		maps.MoveNorthEast()
		break
	case "NW":
		maps.MoveNorthWest()
		break
	case "SE":
		maps.MoveSouthEast()
		break
	case "SW":
		maps.MoveSouthWest()
		break
	case "W":
		maps.MoveWest()
		break
	case "E":
		maps.MoveEast()
		break
	case "N":
		maps.MoveNorth()
		break
	case "S":
		maps.MoveSouth()
		break
	default:
		maps.MoveNorth()
		break
	}
}

/* Movimento ball nella direzione selezionata */
func MoveInDirection(direction string, speed uint8) {
	/* Colore RGB - verde */
	ball.SetRGB(0, 128, 0)

	/* Movimento del robot */
	time.Sleep(wait * time.Millisecond)
	ball.Roll(0, getGradeDirection(direction))
	time.Sleep(wait / 2 * time.Millisecond)
	ball.Roll(speed, getGradeDirection(direction))

	start := time.Now()
	ball.On("collision", func(data interface{}) {

		/* Colore RGB - rosso */
		ball.SetRGB(255, 0, 0)

		/* Tempo di collisione */
		elapsed := time.Since(start)

		/* Metri percorsi */
		//mRide := (elapsed.Seconds() * ms) - elapsed.Seconds()
		mRide := elapsed.Seconds() * ms
		fmt.Printf("Tempo di collisione %f \n", mRide)

		for i := 0; i < interval; i++ {
			setPosition(direction)
		}
		//maps.SetObstacle()
	})
	maps.PrintMap()
}
