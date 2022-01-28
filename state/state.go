package state

import "fmt"

//Declare variables
var HS string
var rev string
var korn string
var kylling string

//Declare positions
var West string = "HS rev korn kylling"
var Boat string = " |_________| "
var East string = ""
var State string = "<W>" + West + Boat + East + "<E> \n"

//var state string = "<W> %s %s %s <E> \n"

//viser status
func ViewState() {
	fmt.Printf(State)
}
