package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	var year = now.Year()
	fmt.Println(year)
}
