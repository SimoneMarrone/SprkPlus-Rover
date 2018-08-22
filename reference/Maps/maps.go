/**
	MAPS  ROBOT
**/

package main

import (
	"fmt"
)

/* Dimensione mappa in quadrati da 10cm  */
const dimensionMap = 25
const dimensionSquare = 10

/* Array mappa con attuale posizione robot */
var maps [dimensionMap][dimensionMap]int
var currentPositionX int
var currentPositionY int

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

	maps[currentPositionX][currentPositionY] = 0
}

/* Check limite mappa */
func checkMapLimit() {

	if currentPositionX > dimensionMap {

		for i := 0; i < dimensionMap; i++ {
			for j := 1; j < dimensionMap; j++ {
				maps[i][j-1] = maps[i][j]
			}
		}
	}

	if currentPositionY > dimensionMap {

		for i := 1; i < dimensionMap; i++ {
			for j := 0; j < dimensionMap; j++ {
				maps[i-1][j] = maps[i][j]
			}
		}
	}

	if currentPositionX < 0 {

	}

	if currentPositionY < 0 {

	}
}

/* Set ostacolo */
func obstacle(obstacleCm int) {

	maps[currentPositionX][currentPositionY] = obstacleCm
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

	for i := 0; i < 5; i++ {

		moveNorth()
		printMap()

		//time.Sleep(500 * time.Millisecond)
	}

	printMap()
}
