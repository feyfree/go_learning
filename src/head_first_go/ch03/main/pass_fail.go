package main

import (
	"fmt"
	"head_first_go/ch03/keyboard"
)

func main() {
	fmt.Println("Enter a grade:")
	float, err := keyboard.GetFloat()
	if err == nil {
		fmt.Println(float)
	}
}
