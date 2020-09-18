package test

import (
	"fmt"
	"log"
	"testing"
)

type ComedyError string

type OverheatError float64

func (o OverheatError) Error() string {
	return fmt.Sprintf("Overheating by %0.2f degrees!", o)
}

func (c ComedyError) Error() string {
	return string(c)
}

func checkTemperature(actual float64, safe float64) error {
	excess := actual - safe
	if excess > 0 {
		return OverheatError(excess)
	}
	return nil
}

func TestCheckTemperature(t *testing.T) {
	err := checkTemperature(121, 100)
	if err != nil {
		log.Fatal(err)
	}
}

func TestComedyError(t *testing.T) {
	var err error
	err = ComedyError("What's a programmer favorite beer? Logger!")
	fmt.Println(err)
}
