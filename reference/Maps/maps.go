/**
	MAPS  ROBOT
**/

package main

import "fmt"

/* Dimensione mappa in quadrati da 10cm  */
const dimensionMap = 25
const dimensionSquare = 10

/* Array mappa con attuale posizione robot */
var maps [dimensionMap][dimensionMap]int
var currentPositionX int
var currentPositionY int

/* Inizializzazione mappa con misure */
func initMap() {

	var i, j int

	for i = 0; i < dimensionMap; i++ {
		for j = 0; j < dimensionMap; j++ {
			maps[i][j] = dimensionSquare
		}
	}
}

/* Set posizione robot */
func setPositiontRobot(x int, y int) {

	currentPositionX = x
	currentPositionY = y

	// DA ELIMINARE
	maps[currentPositionX][currentPositionY] = 100
}

/* Set movimento NORD */
func moveNorth() {

	setPositiontRobot(currentPositionX, currentPositionY-1)
}

/* Set movimento SUD */
func moveSouth() {

	setPositiontRobot(currentPositionX, currentPositionY+1)
}

/* Set movimento EST */
func moveEast() {

	setPositiontRobot(currentPositionX+1, currentPositionY+1)
}

/* Set movimento OVEST */
func moveWest() {

	setPositiontRobot(currentPositionX-1, currentPositionY+1)
}

/* Stampa mappa */
func printMap() {

	var i, j int

	for i = 0; i < dimensionMap; i++ {

		fmt.Printf("\n")

		for j = 0; j < dimensionMap; j++ {
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
	moveNorth()

	printMap()
}
