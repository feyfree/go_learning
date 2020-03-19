package main

import (
	"fmt"
	"os"
)

/**
go run hello_world.go hello
*/
func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello, World", os.Args[1])
	}
	os.Exit(-1)
}
