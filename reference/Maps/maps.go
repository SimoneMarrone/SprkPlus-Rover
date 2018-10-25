/**
	MAPS  ROBOT
**/

package maps

import (
	"fmt"
	"os"
	"os/exec"
)

/* Dimensione mappa  */
const dimensionMap = 25
const square = "="

/* Array mappa con attuale posizione robot */
var maps [dimensionMap][dimensionMap]string
var currentPositionX int
var currentPositionY int

/* Coordinate ostacoli */
type coordObstacle struct {
	north     string
	south     string
	east      string
	west      string
	northWest string
	southWest string
	northEast string
	southEast string
}

/*Inizializzazione mappa con misure*/
func InitMap() {

	for i := 0; i < dimensionMap; i++ {
		for j := 0; j < dimensionMap; j++ {
			maps[i][j] = square
		}
	}

	setPositiontRobot(int(dimensionMap/2), int(dimensionMap/2))
}

/* Set posizione robot */
func setPositiontRobot(x int, y int) {

	currentPositionX = x
	currentPositionY = y
	
	checkMapLimit()
}

/* Check limite mappa */
func checkMapLimit() {

	if currentPositionX > dimensionMap-1 {

		for i := 1; i < dimensionMap; i++ {
			for j := 0; j < dimensionMap; j++ {

				if i == dimensionMap-1 {
					maps[i-1][j] = maps[i][j]
					maps[dimensionMap-1][j] = square
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
					maps[i][dimensionMap-1] = square
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
					maps[0][j] = square
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
					maps[0][j-1] = square
				} else {
					maps[i][j] = maps[i][j-1]
				}
			}
		}
		
		currentPositionY++
	}
}

/* Set ostacolo */
func SetObstacle() {

	maps[currentPositionX][currentPositionY] = "#"
}

/* Get ostacoli su posizione robot */
func LookRound() coordObstacle {

	var tempCurrentPositionX int = currentPositionX
	var tempCurrentPositionY int = currentPositionY

	if currentPositionX <= 0 {

		tempCurrentPositionX++
	}

	if currentPositionY <= 0 {

		tempCurrentPositionY++
	}

	if currentPositionX >= dimensionMap-1 {

		tempCurrentPositionX--
	}

	if currentPositionY >= dimensionMap-1 {
		
		tempCurrentPositionY--
	}

	coord := coordObstacle{
		north:     maps[tempCurrentPositionX-1][tempCurrentPositionY],
		south:     maps[tempCurrentPositionX+1][tempCurrentPositionY],
		east:      maps[tempCurrentPositionX][tempCurrentPositionY+1],
		west:      maps[tempCurrentPositionX][tempCurrentPositionY-1],
		northWest: maps[tempCurrentPositionX-1][tempCurrentPositionY-1],
		southWest: maps[tempCurrentPositionX+1][tempCurrentPositionY-1],
		northEast: maps[tempCurrentPositionX-1][tempCurrentPositionY+1],
		southEast: maps[tempCurrentPositionX+1][tempCurrentPositionY+1]}

	return coord
}

/* Set movimento NORD */
func MoveNorth() {

	setPositiontRobot(currentPositionX-1, currentPositionY)
}

/* Set movimento SUD */
func MoveSouth() {

	setPositiontRobot(currentPositionX+1, currentPositionY)
}

/* Set movimento EST */
func MoveEast() {

	setPositiontRobot(currentPositionX, currentPositionY+1)
}

/* Set movimento OVEST */
func MoveWest() {

	setPositiontRobot(currentPositionX, currentPositionY-1)
}

/* Set movimento NORD-EST */
func MoveNorthEast() {

	setPositiontRobot(currentPositionX-1, currentPositionY+1)
}

/* Set movimento SUD-EST */
func MoveSouthEast() {

	setPositiontRobot(currentPositionX+1, currentPositionY+1)
}

/* Set movimento NORD-OVEST */
func MoveNorthWest() {

	setPositiontRobot(currentPositionX-1, currentPositionY-1)
}

/* Set movimento SUD-OVEST */
func MoveSouthWest() {

	setPositiontRobot(currentPositionX+1, currentPositionY-1)
}

/* Stampa mappa */
func PrintMap() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for i := 0; i < dimensionMap; i++ {

		fmt.Printf("\n")

		for j := 0; j < dimensionMap; j++ {
			fmt.Printf("%s ", maps[i][j])
		}
	}

	fmt.Printf("\n\n## Posizione Robot ##\n")
	fmt.Printf("Posizione X: %d \n", currentPositionX)
	fmt.Printf("Posizione Y: %d \n\n", currentPositionY)
	fmt.Printf("OSTACOLI: %s \n", LookRound())
}
