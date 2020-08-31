package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've choose a random number between 1 to 100, let's guess!")
	fmt.Println(target)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if value == target {
		fmt.Println("You guess right!")
	} else {
		fmt.Println("You guess wrong!")
	}
}
