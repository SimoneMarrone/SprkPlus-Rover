package main

import "fmt"
import "github.com/mndrix/golog"

var test = 1

func main() {

	m := golog.NewMachine()

	fmt.Printf("Primo giro\n")

	m = m.Consult(`blocked(null).`)

	solutions := m.ProveAll(`blocked(X).`)

	for _, solution := range solutions {
		fmt.Printf("%s is busy\n", solution.ByName_("X"))
	}

	fmt.Printf("Secondo giro\n")

	tempM := m.Consult(`blocked(left).`)

	solutions = tempM.ProveAll(`blocked(X).`)
	for _, solution := range solutions {
		fmt.Printf("%s is busy\n", solution.ByName_("X"))
	}

	fmt.Printf("Terzo giro\n")

	tempM = tempM.Consult(`blocked(right).`)

	solutions = tempM.ProveAll(`blocked(X).`)
	for _, solution := range solutions {
		fmt.Printf("%s is busy\n", solution.ByName_("X"))
	}

	fmt.Printf("Quarto giro\n")

	tempM = m

	solutions = tempM.ProveAll(`blocked(X).`)
	for _, solution := range solutions {
		fmt.Printf("%s is busy\n", solution.ByName_("X"))
	}
}
