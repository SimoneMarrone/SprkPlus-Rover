package prolog

import (
	"fmt"
	"time"
	"math/rand"

	"../reference/Maps"
	"github.com/mndrix/golog"
	"../reference/Direction"
)

/* Coordinate ostacoli */
type Direction struct {
	North     map[string]bool
	South     map[string]bool
	East      map[string]bool
	West      map[string]bool
	NorthWest map[string]bool
	SouthWest map[string]bool
	NorthEast map[string]bool
	SouthEast map[string]bool
}

/* All direction are available in this state of Db */
var cleanMachine = golog.NewMachine().Consult("busy(X):-blocked(X). blocked(init).")

/* Machine used for runtime work */
var actualMachine = cleanMachine

/* Machine used for rollback */
var rollbackMachine = cleanMachine

/* Method to assert passed location as busy */
func AssertBusy(b string) {

	/* saving state for rollback */
	rollbackMachine = actualMachine

	actualMachine = actualMachine.Consult("blocked(" + b + ").")
	//fmt.Println("Test direzione busy %t", actualMachine.CanProve("busy("+b+")."))
}

/* Method to erase state */
func Reset() {
	actualMachine = cleanMachine
}

/* Get only free direction */
func FreeDir() []string {
	var freeDir []string

	if !actualMachine.CanProve("busy(w).") {
		freeDir = append(freeDir, "W")
	}
	if !actualMachine.CanProve("busy(e).") {
		freeDir = append(freeDir, "E")
	}
	if !actualMachine.CanProve("busy(n).") {
		freeDir = append(freeDir, "N")
	}
	if !actualMachine.CanProve("busy(nw).") {
		freeDir = append(freeDir, "NW")
	}
	if !actualMachine.CanProve("busy(ne).") {
		freeDir = append(freeDir, "NE")
	}
	if !actualMachine.CanProve("busy(s).") {
		freeDir = append(freeDir, "S")
	}
	if !actualMachine.CanProve("busy(sw).") {
		freeDir = append(freeDir, "SW")
	}
	if !actualMachine.CanProve("busy(se).") {
		freeDir = append(freeDir, "SE")
	}
	return freeDir
}

func SetDirOfMap(){
	
	// N - S - E - O - NO - SO - NE - SE
	//var DIR = [8]string {"N","S","E","O","NO","SO","NE","SE"}
	var DIR Maps.CoordObstacle = Maps.LookAround()
	
	_directions := Direction{
		North :    make(map[string]bool),
		South:     make(map[string]bool),
		East:      make(map[string]bool),
		West:      make(map[string]bool),
		NorthWest: make(map[string]bool),
		SouthWest: make(map[string]bool),
		NorthEast: make(map[string]bool),
		SouthEast: make(map[string]bool)}
		
	for index, b := range DIR.North {
		if(DIR.North[index] == "#"){ _directions.North[string(index-1)] = false } else { _directions.North[string(index-1)] = true }
		if(DIR.South[index] == "#"){ _directions.South[string(index-1)] = false } else { _directions.South[string(index-1)] = true }
		if(DIR.East[index] == "#"){ _directions.East[string(index-1)] = false } else { _directions.East[string(index-1)] = true }
		if(DIR.West[index] == "#"){ _directions.West[string(index-1)] = false } else { _directions.West[string(index-1)] = true }
		if(DIR.NorthWest[index] == "#"){ _directions.NorthWest[string(index-1)] = false } else { _directions.NorthWest[string(index-1)] = true }
		if(DIR.SouthWest[index] == "#"){ _directions.SouthWest[string(index-1)] = false } else { _directions.SouthWest[string(index-1)] = true }
		if(DIR.NorthEast[index] == "#"){ _directions.NorthEast[string(index-1)] = false } else { _directions.NorthEast[string(index-1)] = true }
		if(DIR.SouthEast[index] == "#"){ _directions.SouthEast[string(index-1)] = false } else { _directions.SouthWest[string(index-1)] = true }
		if b == "#" {
			fmt.Println("Ostacolo a distanza: ", index)
		}
	}
		

	fmt.Println(_directions.North["ciao"])

 
	fmt.Println("Direzione nord: ", DIR)
	
	fmt.Println("Direzioni disponibili: ",FreeDir())

	fmt.Println("Sto andando a: ", RandFreeDir())

}

/* Pick a random free direction */

func RandFreeDir() string {
	var x = FreeDir()

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	if(len(x) > 0){
		return x[r.Intn(len(x))]
	}
	fmt.Println("------------------------")
	fmt.Println("No direction available!!")
	fmt.Println("------------------------")
	return "error"

}

func MakeMove(){
	direction.MaxSpeed_Mov(RandFreeDir())
	fmt.Println("Movement finished.")
}

/* Check if an element is present in a list */
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
