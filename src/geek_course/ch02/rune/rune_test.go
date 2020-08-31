package rune

import (
	"fmt"
	"testing"
)

func TestRune1(t *testing.T) {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}

func TestRune2(t *testing.T) {
	var ch = 'A'
	fmt.Print(ch)
	fmt.Print("---separator---")
	var ch01 = '\t'
	fmt.Print(ch01)
	fmt.Print("---END---")
}
