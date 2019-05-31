package prolog

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"

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

var _directions Direction

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
func FreeDir() [][2]string {
	var freeDir [][2]string
	
	if(GetAvailableMov(_directions.North)!=0){ 
		freeDir = append(freeDir,[2]string{"N",strconv.Itoa(GetAvailableMov(_directions.North))})
	}
	if(GetAvailableMov(_directions.NorthEast)!=0){ 
		freeDir = append(freeDir,[2]string{"NE",strconv.Itoa(GetAvailableMov(_directions.NorthEast))})
	}
	if(GetAvailableMov(_directions.NorthWest)!=0){
		freeDir = append(freeDir,[2]string{"NW",strconv.Itoa(GetAvailableMov(_directions.NorthWest))})
	}
	if(GetAvailableMov(_directions.South)!=0){ 
		freeDir = append(freeDir,[2]string{"S",strconv.Itoa(GetAvailableMov(_directions.South))})
	}
	if(GetAvailableMov(_directions.SouthWest)!=0){ 
		freeDir = append(freeDir,[2]string{"SW",strconv.Itoa(GetAvailableMov(_directions.SouthWest))})
	}
	if(GetAvailableMov(_directions.SouthEast)!=0){ 
		freeDir = append(freeDir,[2]string{"SE",strconv.Itoa(GetAvailableMov(_directions.SouthEast))})
	}
	if(GetAvailableMov(_directions.East)!=0){ 
		freeDir = append(freeDir,[2]string{"E",strconv.Itoa(GetAvailableMov(_directions.East))})
	}
	if(GetAvailableMov(_directions.West)!=0){ 
		freeDir = append(freeDir,[2]string{"W",strconv.Itoa(GetAvailableMov(_directions.West))})
	}

	fmt.Println("Prova direzione disponibile: ",strconv.Itoa(GetAvailableMov(_directions.West)))

	return freeDir
}

func SetDirOfMap(){
	
	// N - S - E - O - NO - SO - NE - SE
	//var DIR = [8]string {"N","S","E","O","NO","SO","NE","SE"}
	var DIR Maps.CoordObstacle = Maps.LookAround()
	
	_directions = Direction{
		North :    make(map[string]bool),
		South:     make(map[string]bool),
		East:      make(map[string]bool),
		West:      make(map[string]bool),
		NorthWest: make(map[string]bool),
		SouthWest: make(map[string]bool),
		NorthEast: make(map[string]bool),
		SouthEast: make(map[string]bool)}
		
	for index, _ := range DIR.North {
		var movement_string = "can" + strconv.Itoa(index+1)
		if(index+1 == 1 || index+1 == 3 || index+1 == 6){
			if(DIR.North[index] == "#"){ _directions.North[movement_string] = false } else { _directions.North[movement_string] = true }
			if(DIR.South[index] == "#"){ _directions.South[movement_string] = false } else { _directions.South[movement_string] = true }
			if(DIR.East[index] == "#"){ _directions.East[movement_string] = false } else { _directions.East[movement_string] = true }
			if(DIR.West[index] == "#"){ _directions.West[movement_string] = false } else { _directions.West[movement_string] = true }
			if(DIR.NorthWest[index] == "#"){ _directions.NorthWest[movement_string] = false } else { _directions.NorthWest[movement_string] = true }
			if(DIR.SouthWest[index] == "#"){ _directions.SouthWest[movement_string] = false } else { _directions.SouthWest[movement_string] = true }
			if(DIR.NorthEast[index] == "#"){ _directions.NorthEast[movement_string] = false } else { _directions.NorthEast[movement_string] = true }
			if(DIR.SouthEast[index] == "#"){ _directions.SouthEast[movement_string] = false } else { _directions.SouthWest[movement_string] = true }
		} else{
			continue
		}
	}
		
	fmt.Println("Can 2 ",_directions.North["can2"])
 
	fmt.Println("Direzione nord: ", DIR)
	
	fmt.Println("Direzioni disponibili: ",FreeDir())

	fmt.Println("Sto andando a: ", RandFreeDir())

}

/* Returns max free squares (1-3-6) */
func GetAvailableMov(m map[string]bool) int{
	var movement_type int = 0

	for index, _ := range m {
		switch index {
			case "can1":
				if(m[index]){
					movement_type = 1
				}
			case "can3":
				if(m[index]){
					movement_type = 3
				}
			case "can6":
				if(m[index]){
					movement_type = 6
				}
		}
	}
	return movement_type
}

/* Pick a random free direction */

func RandFreeDir() string {
	var x = FreeDir()

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	_ = r
	_ = x
	/* 
	if(len(x) > 0){
		return x[r.Intn(len(x))][0]
	} */
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
