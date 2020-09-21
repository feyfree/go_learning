package main

import "fmt"

func Socialize() {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello!")
	fmt.Println("How are you today ?")
}

func main() {
	Socialize()
}
