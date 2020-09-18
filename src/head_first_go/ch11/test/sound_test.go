package test

import (
	"fmt"
	"testing"
)

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep, Bop")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs!")
}

func play(n NoiseMaker) {
	n.MakeSound()
}

func TestSoundInterface(t *testing.T) {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	toy.MakeSound()
	toy = Horn("Toyco Blaster")
	toy.MakeSound()
	toy = Robot("SmartRobot")
	toy.MakeSound()
	var noiseMaker NoiseMaker = Robot("Robot")
	// Robot 可以不写    noiseMaker.(Robot)  InterfaceValue.(Asserted Type)
	var robot Robot = noiseMaker.(Robot)
	robot.MakeSound()
	robot.Walk()
	// type assertion 要有容错机制 不然会panic
}
