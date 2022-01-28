package event

import (
	"fmt"

	"github.com/sebmid/crossing-river/state"
)

var position string
var creature string
var creature2 string

//Choose which creature you want to move, and where you want to place them
func PutCreature() {
	fmt.Scan(&position)
	if position == "boat" {
		fmt.Println("Who do you want to move? (HS, rev, kylling, korn)")
		fmt.Scan(&creature)
		if creature == "HS" {
			state.West = "rev korn kylling"
			state.Boat = " |____" + creature + "____| "
			state.East = ""
			state.ViewState()
			fmt.Println("Bring another creature? (Max 2 in one boat)")
			fmt.Scan(&creature2)
			if creature2 == "rev" {
				state.West = "korn kylling"
				state.Boat = " |____" + creature + " " + creature2 + "____| "
				state.East = ""
			} else if creature2 == "korn" {
				state.West = "rev kylling"
				state.Boat = " |____" + creature + " " + creature2 + "____| "
				state.East = ""
			} else if creature2 == "kylling" {
				state.West = "rev korn"
				state.Boat = " |____" + creature + " " + creature2 + "____| "
				state.East = ""
			}
		}
	}
}

//Makes the boat cross to the east
//empties the boat and places them to the east
func CrossRiver() {
	var creature2 string

	fmt.Scan(&creature)
	fmt.Println("Confirm who is the second creature")
	fmt.Scan(&creature2)
	state.Boat = "|_________|"
	state.East = creature + " " + creature2
}
