package main

import (
	"fmt"
	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Println("Enter a grade:")
	float, err := keyboard.GetFloat()
	if err == nil {
		fmt.Println(float)
	}
}
