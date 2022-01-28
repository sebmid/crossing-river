package main

import (
	"fmt"

	"github.com/sebmid/crossing-river/event"
	"github.com/sebmid/crossing-river/state"
)

func main() {
	fmt.Println("Cross the river:")
	state.ViewState()
	event.PutCreature()
	state.ViewState()
}
