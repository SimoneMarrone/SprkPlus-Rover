/**
	MAPS  ROBOT
**/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

/* Dimensione mappa in quadrati da 10cm  */
const dimensionMap = 25
const dimensionSquare = 10

/* Array mappa con attuale posizione robot */
var maps [dimensionMap][dimensionMap]int
var currentPositionX int
var currentPositionY int

/* Coordinate ostacoli */
type coordObstacle struct {
	north     int
	south     int
	east      int
	west      int
	northWest int
	southWest int
	northEast int
	southEast int
}

/* Inizializzazione mappa con misure */
func initMap() {

	for i := 0; i < dimensionMap; i++ {
		for j := 0; j < dimensionMap; j++ {
			maps[i][j] = dimensionSquare
		}
	}
}

/* Set posizione robot */
func setPositiontRobot(x int, y int) {

	currentPositionX = x
	currentPositionY = y

	checkMapLimit()

	//DA CANCELLARE
	maps[currentPositionX][currentPositionY] = 0
}

/* Check limite mappa */
func checkMapLimit() {

	if currentPositionX > dimensionMap-1 {

		for i := 1; i < dimensionMap; i++ {
			for j := 0; j < dimensionMap; j++ {

				if i == dimensionMap-1 {
					maps[i-1][j] = maps[i][j]
					maps[dimensionMap-1][j] = dimensionSquare
				} else {
					maps[i-1][j] = maps[i][j]
				}
			}
		}

		currentPositionX--
	}

	if currentPositionY > dimensionMap-1 {

		for i := 0; i < dimensionMap; i++ {
			for j := 1; j < dimensionMap; j++ {

				if j == dimensionMap-1 {
					maps[i][j-1] = maps[i][j]
					maps[i][dimensionMap-1] = dimensionSquare
				} else {
					maps[i][j-1] = maps[i][j]
				}
			}
		}

		currentPositionY--
	}

	if currentPositionX < 0 {

		for i := dimensionMap - 1; i > 0; i-- {
			for j := dimensionMap - 1; j >= 0; j-- {

				if i == 1 {
					maps[i][j] = maps[i-1][j]
					maps[0][j] = dimensionSquare
				} else {
					maps[i][j] = maps[i-1][j]
				}
			}
		}

		currentPositionX++
	}

	if currentPositionY < 0 {

		for i := dimensionMap - 1; i >= 0; i-- {
			for j := dimensionMap - 1; j > 0; j-- {

				if i == 0 {
					maps[i][j] = maps[i][j-1]
					maps[0][j-1] = dimensionSquare
				} else {
					maps[i][j] = maps[i][j-1]
				}
			}
		}

		currentPositionY++
	}
}

/* Set ostacolo */
func setObstacle(obstacleCm int) {

	maps[currentPositionX][currentPositionY] = obstacleCm
}

/* Get ostacoli su posizione robot */
func getObstacle() coordObstacle {

	coord := coordObstacle{
		north:     maps[currentPositionX-1][currentPositionY],
		south:     maps[currentPositionX+1][currentPositionY],
		east:      maps[currentPositionX][currentPositionY+1],
		west:      maps[currentPositionX][currentPositionY-1],
		northWest: maps[currentPositionX-1][currentPositionY-1],
		southWest: maps[currentPositionX+1][currentPositionY-1],
		northEast: maps[currentPositionX-1][currentPositionY+1],
		southEast: maps[currentPositionX+1][currentPositionY+1]}

	return coord
}

/* Set movimento NORD */
func moveNorth() {

	setPositiontRobot(currentPositionX-1, currentPositionY)
}

/* Set movimento SUD */
func moveSouth() {

	setPositiontRobot(currentPositionX+1, currentPositionY)
}

/* Set movimento EST */
func moveEast() {

	setPositiontRobot(currentPositionX, currentPositionY+1)
}

/* Set movimento OVEST */
func moveWest() {

	setPositiontRobot(currentPositionX, currentPositionY-1)
}

/* Set movimento NORD-EST */
func moveNorthEast() {

	setPositiontRobot(currentPositionX-1, currentPositionY+1)
}

/* Set movimento SUD-EST */
func moveSouthEast() {

	setPositiontRobot(currentPositionX+1, currentPositionY+1)
}

/* Set movimento NORD-OVEST */
func moveNorthWest() {

	setPositiontRobot(currentPositionX-1, currentPositionY-1)
}

/* Set movimento SUD-OVEST */
func moveSouthWest() {

	setPositiontRobot(currentPositionX+1, currentPositionY-1)
}

/* Stampa mappa */
func printMap() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < dimensionMap; i++ {

		fmt.Printf("\n")

		for j := 0; j < dimensionMap; j++ {
			fmt.Printf("%d ", maps[i][j])
		}
	}

	fmt.Printf("\n\n## Posizione Robot ##\n")
	fmt.Printf("Posizione X: %d \n", currentPositionX)
	fmt.Printf("Posizione Y: %d \n\n", currentPositionY)
}

/* MAIN */
func main() {

	/* Inizializzazione mappa */
	initMap()
	setPositiontRobot(int(dimensionMap/2), int(dimensionMap/2))

	for i := 0; i < 10000; i++ {

		switch rand.Intn(8-1) + 1 {
		case 1:
			moveNorth()
			break
		case 2:
			moveSouth()
			break
		case 3:
			moveEast()
			break
		case 4:
			moveWest()
			break
		case 5:
			moveNorthWest()
			break
		case 6:
			moveNorthEast()
			break
		case 7:
			moveSouthWest()
			break
		case 8:
			moveSouthEast()
			break
		}

		printMap()
		fmt.Printf("OSTACOLI: %d \n", getObstacle())

		time.Sleep(200 * time.Millisecond)
	}
}
