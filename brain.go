package main

import (
	"fmt"

	"github.com/mndrix/golog"
)

/* All direction are available in this state of DB */
var cleanMachine = golog.NewMachine().Consult("busy(init). busy(L). blocked(X):-busy(X). free(X):- \\+busy(X).")

/* Machine used for runtime work */
var actualMachine = cleanMachine

/* Machine used for rollback */
var rollbackMachine = cleanMachine

/* Method to check if direction is busy */
func checkDirection(x string) {

	var solutions = actualMachine.ProveAll("blocked(X).")
	var solutionList []string

	for _, solution := range solutions {

		var s = solution.ByName_("X").String()
		solutionList = append(solutionList, s)
	}

	if stringInSlice(x, solutionList) {
		fmt.Println(x + " is busy")
	} else {
		fmt.Println(x + " is free")
	}

}

/* Method to assert passed location as busy */
func assertBusy(b string) {

	/* saving state for rollback */
	rollbackMachine = actualMachine

	actualMachine = actualMachine.Consult("busy(" + b + ").")
}

/* Method to erase state */
func reset() {
	actualMachine = cleanMachine
}

/* Get only free direction */
func freeDir(x string) {
	fmt.Println(actualMachine.CanProve("free(" + x + ")."))
}

/* Main method (test) */

func main() {

	/* 	checkDirection("L")
	   	checkDirection("R")
	   	checkDirection("F")
	   	checkDirection("FL")
	   	checkDirection("FR")
	   	checkDirection("B")
	   	checkDirection("BL")
	   	checkDirection("BR")

	   	assertBusy("F")

	   	checkDirection("L")
	   	checkDirection("R")
	   	checkDirection("F")
	   	checkDirection("FL")
	   	checkDirection("FR")
	   	checkDirection("B")
	   	checkDirection("BL")
	   	checkDirection("BR") */

	freeDir("F")

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
