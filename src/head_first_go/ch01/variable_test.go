package ch01

import (
	"fmt"
	"testing"
)

func TestVariableAssignment(t *testing.T) {
	var number = 4
	// var number int = 4
	fmt.Println(number)

	var length, width float64 = 1.2, 3.4
	fmt.Println(length, width)

	quantity := 4
	fmt.Println(quantity)

	customerName := "feyfree"
	fmt.Println(customerName)
}

func TestConversion(t *testing.T) {
	var length float64 = 1.2
	var width int = 2
	// type mismatch
	//fmt.Println("Area is", length*width)
	//fmt.Println("length > width ?", length > width)

	fmt.Println("Area is", length*float64(width))
	fmt.Println("length > width ?", length > float64(width))

	fmt.Println(int64(length))

}
