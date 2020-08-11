package ch02

import (
	"fmt"
	"strings"
	"testing"
)

func TestCalling(t *testing.T) {
	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
