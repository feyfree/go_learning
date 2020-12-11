package test

import (
	"fmt"
	"testing"
)

type MyType string
type MyNumber int

func (n *MyNumber) Display() { fmt.Println(*n) }
func (n *MyNumber) Double()  { *n *= 2 }

func (m MyType) method() {
	fmt.Println("Method with value receiver!")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer value receiver!")
}

func TestPointerMethod(t *testing.T) {
	value := MyType("A value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
	MyType("Another value").method()
	// cannot call instead you need to use a variable to store,
	// go will do the my_conversion_test for you
	//MyType("Another value").pointerMethod()
}

func TestMyNumber(t *testing.T) {
	number := MyNumber(4)
	number.Double()
	number.Display()
}
