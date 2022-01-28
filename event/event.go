package event

import (
	"fmt"

	"github.com/sebmid/crossing-river/state"
)

var position string
var creature string

//Choose which creature you want to move, and where you want to place them
func PutCreature() {
	fmt.Scan(&position)
	if position == "boat" {
		fmt.Println("Who do you want to move? (HS, rev, kylling, korn)")
		fmt.Scan(&creature)
		if creature == "HS" {
			state.West = "rev korn kylling"
			state.Boat = " |____HS____| "
			state.East = ""
			state.ViewState()
			fmt.Println("Bring another creature? (Max 2 in one boat)")
			fmt.Scan(&creature)
			if creature == "rev" {
				state.West = "korn kylling"
				state.Boat = " |___rev_HS___| "
				state.East = ""
				state.ViewState()
			} else if creature == "korn" {
				state.West = "rev kylling"
				state.Boat = " |___korn_HS___| "
				state.East = ""
				state.ViewState()
			} else if creature == "kylling" {
				state.West = "rev korn"
				state.Boat = " |___kylling_HS___| "
				state.East = ""
				state.ViewState()
			}
		}
	}
}

//Makes the boat cross to the east
//empties the boat and places them to the east
func CrossRiver() {
	var creature2 string

	fmt.Scan(&creature)
	fmt.Println("What is the second creature?")
	fmt.Scan(&creature2)
	state.Boat = "|_________|"
	state.East = creature + " " + creature2
	state.ViewState()
}
