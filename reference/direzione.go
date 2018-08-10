package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// SOLO PER CONTROLLO - DA ELIMINARE - SCELTA DIREZIONE
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Direzionare verso :  (F,B,L,R)")
	text, _ := reader.ReadString('\n')

	if strings.TrimRight(text, "\n") == "B" {
		directionBack()
	}
	if strings.TrimRight(text, "\n") == "F" {
		directionForward()
	}
	if strings.TrimRight(text, "\n") == "L" {
		directionLeft()
	}
	if strings.TrimRight(text, "\n") == "R" {
		directionRight()
	}
	// SOLO PER CONTROLLO - DA ELIMINARE - SCELTA DIREZIONE PIù CAMMINO VERSO QUELLA DIREZIONE (impostata velocità fissa a 5)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Andare verso : (F,B,L,R)")
	text, _ := reader.ReadString('\n')

	if strings.TrimRight(text, "\n") == "B" {
		directionBack()
	}
	if strings.TrimRight(text, "\n") == "F" {
		directionForward()
	}
	if strings.TrimRight(text, "\n") == "L" {
		directionLeft()
	}
	if strings.TrimRight(text, "\n") == "R" {
		directionRight()
	}
}

// FUNZIONI PER FAR GIRARE SPHERO VERSO UNA DIREZIONE
func directionBack() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FA GIRARE SPHERO NEL SENSO OPPOSTO RISPETTO A QUELLO IN CUI SI TROVA
	ball.Roll(0, 180)
}
func directionForward() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)

	ball.Roll(0, 0)
}
func directionLeft() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FA GIRARE SPHERO VERSO SX DAL PUNTO IN CUI SI TROVA
	ball.Roll(0, 270)
}
func directionRight() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FA GIRARE SPHERO VERSO DX DAL PUNTO IN CUI SI TROVA
	ball.Roll(0, 90)
}

// FUNZIONI PER FAR ANDARE SPHERO VERSO UNA DIREZIONE
func goBack() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FAR ANDARE SPHERO NEL SENSO OPPOSTO RISPETTO A QUELLO IN CUI SI TROVA
	ball.Roll(0, 180)
	ball.Roll(5, 0)
}
func goForward() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FAR ANDARE SPHERO NEL SENSO DRITTO NEL SENSO IN CUI SI TROVA
	ball.Roll(0, 0)
	ball.Roll(5, 0)
}
func goRight() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FAR ANDARE SPHERO VERSO DESTRA DAL PUNTO IN CUI SI TROVA
	ball.Roll(0, 90)
	ball.Roll(5, 0)
}
func goLeft() {
	bleAdaptor := ble.NewClientAdaptor(os.Args[1])
	ball := sprkplus.NewDriver(bleAdaptor)
	// FAR ANDARE SPHERO VERSO SINISTRA DAL PUNTO IN CUI SI TROVA
	ball.Roll(0, 270)
	ball.Roll(5, 0)
}
