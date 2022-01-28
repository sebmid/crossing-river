package state

import (
	"testing"
)

func TestViewState(t *testing.T) {
	wanted := "<W>HS rev korn kylling |_________| <E> \n"
	got := State
	if got != wanted {
		t.Errorf("You got:%s but we want it to be:%s", got, wanted)
	}
}
