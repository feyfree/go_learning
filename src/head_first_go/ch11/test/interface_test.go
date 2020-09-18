package test

import (
	"fmt"
	"testing"
)

type Switch string

func (s *Switch) toggle() {
	if *s == "On" {
		*s = "Off"
	} else {
		*s = "On"
	}
	fmt.Println(*s)
}

type Toggleable interface {
	toggle()
}

func TestInterface(t *testing.T) {
	s := Switch("Off")
	// When Go decides whether a value satisfies an interface,
	//pointer methods aren’t included for direct values. But
	//they are included for pointers. So the solution is to
	//assign a pointer to a Switch to the Toggleable variable,
	//instead of a direct Switch value
	var to Toggleable = &s
	s.toggle()
	to.toggle()
}
