package test

import (
	"fmt"
	"testing"
)

func one() {
	defer fmt.Println("defer in one()")
	two()
}

func two() {
	defer fmt.Println("defer in two()")
	three()
}

func three() {
	defer fmt.Println("defer in three()")
	panic("See what happened!")
}

func TestStack(t *testing.T) {
	one()
}
