package prolog

import (
	"fmt"
	"math/rand"
	"time"

	"../reference/Maps"
	"github.com/mndrix/golog"
)

/* All direction are available in this state of Db */
var cleanMachine = golog.NewMachine().Consult("busy(init). blocked(X):-busy(X). free(X):- \\+blocked(X).")

/* Machine used for runtime work */
var actualMachine = cleanMachine

/* Machine used for rollback */
var rollbackMachine = cleanMachine

/* Method to check if direction is busy */
func CheckDirection(x string) bool{

	var solutions = actualMachine.ProveAll("blocked(X).")
	var solutionlist []string

	for _, solution := range solutions {

		var s = solution.ByName_("X").String()
		solutionlist = append(solutionlist, s)
	}

	if stringInSlice(x, solutionlist) {
		fmt.Println(x + " is busy")
		return false
	} else {
		fmt.Println(x + " is free")
		return true
	}

}

/* Method to assert passed location as busy */
func AssertBusy(b string) {

	/* saving state for rollback */
	rollbackMachine = actualMachine

	actualMachine = actualMachine.Consult("busy(" + b + ").")
	//fmt.Println("Test direzione busy %t", actualMachine.CanProve("busy("+b+")."))
}

/* Method to erase state */
func Reset() {
	actualMachine = cleanMachine
}

/* Get only free direction */
func FreeDir() []string {
	var freeDir []string

	if actualMachine.CanProve("free(w).") {
		freeDir = append(freeDir, "W")
	}
	if actualMachine.CanProve("free(e).") {
		freeDir = append(freeDir, "E")
	}
	if actualMachine.CanProve("free(n).") {
		freeDir = append(freeDir, "N")
	}
	if actualMachine.CanProve("free(nw).") {
		freeDir = append(freeDir, "NW")
	}
	if actualMachine.CanProve("free(ne).") {
		freeDir = append(freeDir, "NE")
	}
	if actualMachine.CanProve("free(s).") {
		freeDir = append(freeDir, "S")
	}
	if actualMachine.CanProve("free(sw).") {
		freeDir = append(freeDir, "SW")
	}
	if actualMachine.CanProve("free(se).") {
		freeDir = append(freeDir, "SE")
	}
	return freeDir	//return freeDir
}

func SetDirOfMap(){
	
	// N - S - E - O - NO - SO - NE - SE
	//var DIR = [8]string {"N","S","E","O","NO","SO","NE","SE"}
	
	var DIR maps.CoordObstacle = maps.LookRound()

	if(DIR.North == "#") { AssertBusy("n") }
	if(DIR.NorthEast == "#") { AssertBusy("ne") }
	if(DIR.NorthWest == "#") { AssertBusy("nw") }
	if(DIR.South == "#") { AssertBusy("s") }
	if(DIR.SouthEast == "#") { AssertBusy("se") }
	if(DIR.SouthWest == "#") { AssertBusy("sw") }
	if(DIR.West == "#") { AssertBusy("w") }
	if(DIR.East == "=") { AssertBusy("e") }
	
	fmt.Println(actualMachine.CanProve("free(e)."))
	AssertBusy("e")
	fmt.Println(actualMachine.CanProve("busy(e)."))
	fmt.Println(actualMachine.CanProve("free(e)."))
}

/* Pick a random free direction */

func RandFreeDir() string {
	var x = FreeDir()

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	return x[r.Intn(len(x))]
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
