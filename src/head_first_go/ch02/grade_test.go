package ch02

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestGrade(t *testing.T) {
	fmt.Println("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
