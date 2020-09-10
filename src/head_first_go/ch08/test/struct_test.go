package test

import (
	"fmt"
	"testing"
)

var myStruct struct {
	number float64
	word   string
	toggle bool
}

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
	engine   string
}

func showInfo(p car) {
	fmt.Print("name:"+p.name, p.topSpeed)
}

func TestStruct1(t *testing.T) {
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println(myStruct)
}

func (p *car) Engine() {
	fmt.Println(p.engine)
}

func TestStruct2(t *testing.T) {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 250.0
	porsche.engine = "wuwu"
	fmt.Println(porsche)
	showInfo(porsche)
	porsche.Engine()

}
