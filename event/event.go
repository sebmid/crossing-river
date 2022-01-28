package event

import (
	"fmt"

	"github.com/sebmid/crossing-river/state"
)

var position string
var creature string

//Where do you want do put the homosapien
func PutCreature() {
	fmt.Println("Where do you want to move the creature? (west, boat, east)")
	fmt.Scan(&position)
	if position == "west" {
		fmt.Println("Who do you want to move? (HS, rev, kylling, korn)")
		fmt.Scan(&creature)
		if creature == "HS" {
			state.West = "HS rev korn kylling"
		}
	} else if position == "boat" {
		fmt.Println("Who do you want to move? (HS, rev, kylling, korn)")
		if creature == "HS" {
			state.Boat = " |____HS____| "
			state.West = "rev korn kylling"
			state.East = ""
		}
	}
}
