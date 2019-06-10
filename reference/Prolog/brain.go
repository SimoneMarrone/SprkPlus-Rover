package prolog

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"../Direction"
	"../Maps"
	"github.com/mndrix/golog"
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
var cleanMachine = golog.NewMachine().Consult(
	"available(X,Y):- inited(X,Y), \\+ blocked(X,Y). " + // inited and not blocked
		"inited(n,1)." +
		"inited(n,3)." +
		"inited(n,6)." +
		"inited(s,1)." +
		"inited(s,3)." +
		"inited(s,6)." +
		"inited(w,1)." +
		"inited(w,3)." +
		"inited(w,6)." +
		"inited(e,1)." +
		"inited(e,3)." +
		"inited(e,6)." +
		"inited(nw,1)." +
		"inited(nw,3)." +
		"inited(nw,6)." +
		"inited(ne,1)." +
		"inited(ne,3)." +
		"inited(ne,6)." +
		"inited(sw,1)." +
		"inited(sw,3)." +
		"inited(sw,6)." +
		"inited(se,1)." +
		"inited(se,3)." +
		"inited(se,6)." +
		"blocked(init,0).")

/* Machine used for runtime work */
var actualMachine = cleanMachine

/* Machine used for rollback */
var rollbackMachine = cleanMachine

var _directions Direction

/* Method to assert passed location as busy --> Hit obstacle */
func AssertBusy(direction string, distance string) {

	/* saving state for rollback */
	rollbackMachine = actualMachine

	actualMachine = actualMachine.Consult("blocked(" + direction + "," + distance + ").")
	//fmt.Println("Test direzione busy %t", actualMachine.CanProve("busy("+b+")."))
}

/* Method to erase state */
func Reset() {
	actualMachine = cleanMachine
}

/* Get only free direction */
func FreeDir() map[string]int {

	_dir_map := make(map[string]int)

	_directions := actualMachine.ProveAll("available(X,Y).")

	for _, solution := range _directions {

		var _dir_key string = fmt.Sprintf("%s", solution.ByName_("X"))
		_dist_available, _ := strconv.Atoi(fmt.Sprintf("%s", solution.ByName_("Y")))

		//fmt.Printf("%s can move by %s cells.\n", solution.ByName_("X"), solution.ByName_("Y"))

		if _dir_map[_dir_key] < _dist_available {
			_dir_map[_dir_key] = _dist_available
		}
	}

	return _dir_map
}

func SetDirOfMap() {

	// N - S - E - O - NO - SO - NE - SE
	//var DIR = [8]string {"N","S","E","O","NO","SO","NE","SE"}
	var DIR Maps.CoordObstacle = Maps.LookAround()

	_directions = Direction{
		North:     make(map[string]bool),
		South:     make(map[string]bool),
		East:      make(map[string]bool),
		West:      make(map[string]bool),
		NorthWest: make(map[string]bool),
		SouthWest: make(map[string]bool),
		NorthEast: make(map[string]bool),
		SouthEast: make(map[string]bool)}

	for index, _ := range DIR.North {
		var movement_string = "can" + strconv.Itoa(1)
		if index+1 == 1 || index+1 == 3 || index+1 == 6 {
			if DIR.North[index] == "#" {
				_directions.North[movement_string] = false
				AssertBusy("n",strconv.Itoa(1))
				AssertBusy("n",strconv.Itoa(3))
				AssertBusy("n",strconv.Itoa(6))

				AssertBusy("nw",strconv.Itoa(1))
				AssertBusy("nw",strconv.Itoa(3))
				AssertBusy("nw",strconv.Itoa(6))

				AssertBusy("ne",strconv.Itoa(1))
				AssertBusy("nw",strconv.Itoa(3))
				AssertBusy("nw",strconv.Itoa(6))
			} else {
				_directions.North[movement_string] = true
			}
			if DIR.South[index] == "#" {
				_directions.South[movement_string] = false
				AssertBusy("s",strconv.Itoa(1))
				AssertBusy("s",strconv.Itoa(3))
				AssertBusy("s",strconv.Itoa(6))

				AssertBusy("sw",strconv.Itoa(1))
				AssertBusy("sw",strconv.Itoa(3))
				AssertBusy("sw",strconv.Itoa(6))
				
				AssertBusy("se",strconv.Itoa(1))
				AssertBusy("se",strconv.Itoa(3))
				AssertBusy("se",strconv.Itoa(6))
			} else {
				_directions.South[movement_string] = true
			}
			if DIR.East[index] == "#" {
				_directions.East[movement_string] = false
				AssertBusy("e",strconv.Itoa(1))
				AssertBusy("e",strconv.Itoa(3))
				AssertBusy("e",strconv.Itoa(6))

				AssertBusy("se",strconv.Itoa(1))
				AssertBusy("se",strconv.Itoa(3))
				AssertBusy("se",strconv.Itoa(6))
				
				AssertBusy("ne",strconv.Itoa(1))
				AssertBusy("ne",strconv.Itoa(3))
				AssertBusy("ne",strconv.Itoa(6))
			} else {
				_directions.East[movement_string] = true
			}
			if DIR.West[index] == "#" {
				_directions.West[movement_string] = false
				AssertBusy("w",strconv.Itoa(1))
				AssertBusy("w",strconv.Itoa(3))
				AssertBusy("w",strconv.Itoa(6))

				AssertBusy("nw",strconv.Itoa(1))
				AssertBusy("nw",strconv.Itoa(3))
				AssertBusy("nw",strconv.Itoa(6))

				AssertBusy("sw",strconv.Itoa(1))
				AssertBusy("sw",strconv.Itoa(3))
				AssertBusy("sw",strconv.Itoa(6))
			} else {
				_directions.West[movement_string] = true
			}
			if DIR.NorthWest[index] == "#" {
				_directions.NorthWest[movement_string] = false
				AssertBusy("nw",strconv.Itoa(1))
				AssertBusy("nw",strconv.Itoa(3))
				AssertBusy("nw",strconv.Itoa(6))

				AssertBusy("n",strconv.Itoa(1))
				AssertBusy("n",strconv.Itoa(3))
				AssertBusy("n",strconv.Itoa(6))

				AssertBusy("w",strconv.Itoa(1))
				AssertBusy("w",strconv.Itoa(3))
				AssertBusy("w",strconv.Itoa(6))
			} else {
				_directions.NorthWest[movement_string] = true
			}
			if DIR.SouthWest[index] == "#" {
				_directions.SouthWest[movement_string] = false
				AssertBusy("sw",strconv.Itoa(1))
				AssertBusy("sw",strconv.Itoa(3))
				AssertBusy("sw",strconv.Itoa(6))

				AssertBusy("s",strconv.Itoa(1))
				AssertBusy("s",strconv.Itoa(3))
				AssertBusy("s",strconv.Itoa(6))

				AssertBusy("w",strconv.Itoa(1))
				AssertBusy("w",strconv.Itoa(3))
				AssertBusy("w",strconv.Itoa(6))
			} else {
				_directions.SouthWest[movement_string] = true
			}
			if DIR.NorthEast[index] == "#" {
				_directions.NorthEast[movement_string] = false
				AssertBusy("ne",strconv.Itoa(1))
				AssertBusy("ne",strconv.Itoa(3))
				AssertBusy("ne",strconv.Itoa(6))

				AssertBusy("n",strconv.Itoa(1))
				AssertBusy("n",strconv.Itoa(3))
				AssertBusy("n",strconv.Itoa(6))

				AssertBusy("e",strconv.Itoa(1))
				AssertBusy("e",strconv.Itoa(3))
				AssertBusy("e",strconv.Itoa(6))
			} else {
				_directions.NorthEast[movement_string] = true
			}
			if DIR.SouthEast[index] == "#" {
				_directions.SouthEast[movement_string] = false
				AssertBusy("se",strconv.Itoa(1))
				AssertBusy("se",strconv.Itoa(3))
				AssertBusy("se",strconv.Itoa(6))

				AssertBusy("s",strconv.Itoa(1))
				AssertBusy("s",strconv.Itoa(3))
				AssertBusy("s",strconv.Itoa(6))

				AssertBusy("e",strconv.Itoa(1))
				AssertBusy("e",strconv.Itoa(3))
				AssertBusy("e",strconv.Itoa(6))
			} else {
				_directions.SouthEast[movement_string] = true
			}
		} else {
			continue
		}
	}
}

/* Pick a random free direction */

func RandFreeDir() (string, int) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	var x = FreeDir()

	fmt.Println("Log mappa: ", x)

	if len(x) > 0 {

		keys := reflect.ValueOf(x).MapKeys()
		strkeys := make([]string, len(keys))

		for i := 0; i < len(keys); i++ {
			strkeys[i] = keys[i].String()
		}

		key := strkeys[r.Intn(len(x))]

		return key, x[key]
	}

	fmt.Println("------------------------")
	fmt.Println("No direction available!!")
	fmt.Println("------------------------")
	return "error", 0
}

func MakeMove() (int, string) {
	d, s := RandFreeDir()

	switch s {
	case 1:
		direction.LowSpeed_Mov(d)
		break
	case 3:
		direction.MidSpeed_Mov(d)
		break
	case 6:
		direction.MaxSpeed_Mov(d)
		break
	}

	fmt.Println("DIRECTION: ", d)
	fmt.Println("SPEED: ", s)

	return s, d
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
