package main

import (
	"fmt"

	"github.com/sebmid/crossing-river/event"
	"github.com/sebmid/crossing-river/state"
)

func main() {
	fmt.Println("Cross the river:")
	state.ViewState()
	fmt.Println("Where do you want to move one of the creatures? (west, boat, east)")
	event.PutCreature()
	fmt.Println("Who is the first creature crossing the river?")
	event.CrossRiver()
}
