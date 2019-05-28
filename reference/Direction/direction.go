package direction

import (
	"fmt"
	"time"

	"../Maps"
	"gobot.io/x/gobot/platforms/sphero/ollie"
)

const ms float64 = 0.4347826
const wait time.Duration = 2500
const speed uint8 = 25
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
func MoveInDirection(direction string) {

	/* Imposta la stabilizzazione */
	ball.SetStabilization(true)

	/* Movimento del robot */
	time.Sleep(wait * time.Millisecond)
	ball.Roll(0, getGradeDirection(direction))
	time.Sleep(wait / 2 * time.Millisecond)
	ball.Roll(speed, getGradeDirection(direction))

	start := time.Now()
	var isCollision = false
	ball.On("collision", func(data interface{}) {
		isCollision = true

		/* Tempo di collisione */
		elapsed := time.Since(start)
		elapsed.Seconds()

		/* Metri percorsi */
		//mRide := (elapsed.Seconds() * ms) - elapsed.Seconds()
		//fmt.Printf("Percorso %f \n", mRide-(elapsed.Seconds()))
		fmt.Printf("COLLISIONE")

		for i := 0; i < interval; i++ {
			setPosition(direction)
		}
		maps.SetObstacle()
	})

	if !isCollision {
		fmt.Printf("NESSUNA COLLISIONE")
		for i := 0; i < interval; i++ {
			setPosition(direction)
		}
	}
	maps.PrintMap()
}
