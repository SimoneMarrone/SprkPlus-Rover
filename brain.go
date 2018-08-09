package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mndrix/golog"
)

/* All direction are available in this state of Db */
var cleanMachine = golog.NewMachine().Consult("busy(init). blocked(X):-busy(X). free(X):- \\+blocked(X).")

/* Machine used for runtime work */
var actualMachine = cleanMachine

/* Machine used for rollback */
var rollbackMachine = cleanMachine

/* Method to check if direction is busy */
func checkDirection(x string) {

	var solutions = actualMachine.ProveAll("blocked(X).")
	var solutionlist []string

	for _, solution := range solutions {

		var s = solution.ByName_("X").String()
		solutionlist = append(solutionlist, s)
	}

	if stringInSlice(x, solutionlist) {
		fmt.Println(x + " is busy")
	} else {
		fmt.Println(x + " is free")
	}

}

/* Method to assert passed location as busy */
func assertbusy(b string) {

	/* saving state for rollback */
	rollbackMachine = actualMachine

	actualMachine = actualMachine.Consult("busy(" + b + ").")
}

/* Method to erase state */
func reset() {
	actualMachine = cleanMachine
}

/* Get only free direction */
func freeDir() []string {
	var freeDir []string

	if actualMachine.CanProve("free(l).") {
		freeDir = append(freeDir, "L")
	}
	if actualMachine.CanProve("free(r).") {
		freeDir = append(freeDir, "R")
	}
	if actualMachine.CanProve("free(f).") {
		freeDir = append(freeDir, "F")
	}
	if actualMachine.CanProve("free(fl).") {
		freeDir = append(freeDir, "FL")
	}
	if actualMachine.CanProve("free(fr).") {
		freeDir = append(freeDir, "FR")
	}
	if actualMachine.CanProve("free(b).") {
		freeDir = append(freeDir, "B")
	}
	if actualMachine.CanProve("free(bl).") {
		freeDir = append(freeDir, "BL")
	}
	if actualMachine.CanProve("free(br).") {
		freeDir = append(freeDir, "BR")
	}

	return freeDir
}

/* Pick a random free direction */

func randFreeDir() string {
	var x = freeDir()

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	return x[r.Intn(len(x))]
}

/* Main method (test) */

func main() {

	/* 	checkDirection("l")
	   	checkDirection("r")
	   	checkDirection("f")
	   	checkDirection("fl")
	   	checkDirection("fr")
	   	checkDirection("b")
	   	checkDirection("bl")
	   	checkDirection("br")

	   	assertbusy("f")

	   	checkDirection("l")
	   	checkDirection("r")
	   	checkDirection("f")
	   	checkDirection("fl")
	   	checkDirection("fr")
	   	checkDirection("b")
	   	checkDirection("bl")
	   	checkDirection("br") */
	/*
		fmt.Println(actualMachine.CanProve("free(l).")) */

	fmt.Println(freeDir())
	assertbusy("l")
	assertbusy("br")
	fmt.Println(freeDir())
	fmt.Println(randFreeDir())

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
