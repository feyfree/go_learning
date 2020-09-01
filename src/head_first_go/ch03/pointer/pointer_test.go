package pointer

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPointer1(t *testing.T) {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)
}

func TestReflect(t *testing.T) {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))
	fmt.Println(reflect.TypeOf(myInt))
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))
	fmt.Println(reflect.TypeOf(myFloat))

	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)
	fmt.Println(reflect.TypeOf(myBoolPointer))
}

func TestUpdateByPointer(t *testing.T) {
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}

func createPointer() *float64 {
	var myFloat = 1.0
	return &myFloat
}

func TestPointerPassing(t *testing.T) {
	pointer := createPointer()
	fmt.Println(pointer)
}

func printBoolPointer(pointer *bool) {
	fmt.Println(*pointer)
}

func TestPointerParameter(t *testing.T) {
	var myBool = true
	printBoolPointer(&myBool)
}
