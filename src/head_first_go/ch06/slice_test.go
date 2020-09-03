package ch06

import (
	"fmt"
	"testing"
)

func TestSliceInit(t *testing.T) {
	// declare a slice 声明
	var notes []string
	// create a slice 创建
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	newNotes := make([]string, 7)
	fmt.Println(len(newNotes))
}

func TestSliceLiteral(t *testing.T) {
	notes := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	primes := []int{2, 3, 5}
	fmt.Println(primes[0], primes[1], primes[2])
}
