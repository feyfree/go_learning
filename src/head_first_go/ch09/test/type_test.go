package test

import (
	"fmt"
	"testing"
)

type Liters float64
type Gallons float64
type Title string
type Message string
type Number int

func (m Message) ShowMessage() {
	fmt.Println("This is a message : ", m)
}

func TestMethod(t *testing.T) {
	value := Message("Hi")
	value.ShowMessage()
}

func (n *Number) Double() {
	*n *= 2
}

func TestPointer(t *testing.T) {
	number := Number(4)
	number.Double()
	fmt.Println(number)
}

func TestType(t *testing.T) {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(carFuel, busFuel)
	carFuel = Gallons(Liters(40) * 0.264)
	busFuel = Liters(Gallons(63) * 3.785)
}

func TestOperator(t *testing.T) {
	fmt.Println(Gallons(10) * Gallons(100))
	fmt.Println(Title("hi") == Title("hi"))
}

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func TestConvertFunc(t *testing.T) {
	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(10.0)
	busFuel = Liters(240.0)
	fmt.Println(ToGallons(busFuel), ToLiters(carFuel))
}
