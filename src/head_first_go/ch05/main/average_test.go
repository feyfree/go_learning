package main

import (
	"fmt"
	"head_first_go/ch05/datafile"
	"log"
	"testing"
)

// 一个package 只能有一个main 注意
func TestAverage(t *testing.T) {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f \n", sum/sampleCount)
}
