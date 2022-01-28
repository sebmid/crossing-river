package state

import "fmt"

//Declare variables
//Declare positions
var West string = "HS rev korn kylling"
var Boat string = " |_________| "
var East string = ""

//var State string = "<W>" + West + Boat + East + "<E> \n"

var State string = "<W> %s %s %s <E> \n"

//viser status
func ViewState() {
	fmt.Printf(State, West, Boat, East)
}
