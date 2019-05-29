package direction

import (
	"fmt"
	"time"

	"gobot.io/x/gobot/platforms/sphero/ollie"
	"../Maps"
)
const ms float64 = 0.4347826
var ball *ollie.Driver

var wait time.Duration = 1000
var wait_2 time.Duration = 500
var speed uint8 = 50

func SetBall(main_ball *ollie.Driver) {
	ball = main_ball
}

func DirectionNE() {

	/* Movimento del robot */
	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn NE")
	ball.Roll(0, 45)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 45)
	

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso NE! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveNorthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})
	}
func DirectionNO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn NO")
	ball.Roll(0, 315)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 315)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso NO! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveNorthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}
func DirectionSE() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn SE")
	ball.Roll(0, 135)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 135)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso SE! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveSouthEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}
func DirectionSO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn SO")
	ball.Roll(0, 225)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 225)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso SO! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveSouthWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}
func DirectionO() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn left")
	ball.Roll(0, 270)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 270)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso O! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveWest()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}
func DirectionE() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn right")
	ball.Roll(0, 90)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 90)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso E! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveEast()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}

func DirectionN() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn forward")
	ball.Roll(0, 0)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 0)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso N! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveNorth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})
}
func DirectionS() {

	time.Sleep(wait * time.Millisecond)
	fmt.Print("\n turn back")
	ball.Roll(0, 180)
	time.Sleep(wait_2 * time.Millisecond)
	ball.Roll(speed, 180)

	ball.On("collision", func(data interface{}) {
		start := time.Now()
		fmt.Print("\n Ostacolo trovato verso S! \n")
		//Get time on collision in second
		elapsed := time.Since(start)
		elapsed.Seconds()
		//Get meter travel
		mRide := elapsed.Seconds() * ms
		fmt.Printf("raw data: %f \n", elapsed.Seconds())
		if elapsed.Seconds() <= 1.15 {
			fmt.Printf("elapsed <30cm = %f \n", mRide-(elapsed.Seconds()*0.165))

			/* Movimento del robot sulla mappa */
			for i := 1; i <= 6; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.15 <= elapsed.Seconds() && elapsed.Seconds() <= 1.40 {
			fmt.Printf("elapsed ca40cm= %f \n", mRide-(elapsed.Seconds()*0.110))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 8; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.40 < elapsed.Seconds() && elapsed.Seconds() <= 1.50 {
			fmt.Printf("elapsed +40cm= %f \n", mRide-(elapsed.Seconds()*0.085))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 10; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.50 < elapsed.Seconds() && elapsed.Seconds() <= 1.75 {
			fmt.Printf("elapsed 60cm= %f \n", mRide-(elapsed.Seconds()*0.092))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 12; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.70 < elapsed.Seconds() && elapsed.Seconds() <= 1.99 {
			fmt.Printf("elapsed 70cm= %f \n", mRide-(elapsed.Seconds()*0.07))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 14; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else if 1.99 < elapsed.Seconds() && elapsed.Seconds() <= 2.40 {
			fmt.Printf("elapsed 80-90cm = %f \n", mRide-(elapsed.Seconds()*0.06))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 17; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		} else {
			fmt.Printf("elapsed 90+cm = %f \n", mRide-(elapsed.Seconds()*0.054))
			/* Movimento del robot sulla mappa */
			for i := 1; i <= 20; i++ {
				maps.MoveSouth()
			}
			/* SEGNALA OSTACOLO SULLA MAPPA */
				maps.SetObstacle()
				maps.PrintMap()
		}
		
	})

}
