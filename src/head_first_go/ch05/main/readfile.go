package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// "./"表示是src 目录下面
	_, _ = os.Create("./hi.txt")
	file, err := os.Open("src\\head_first_go\\ch05\\main\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

}
