package test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	switch v := p.(type) {
	case float64:
		fmt.Println("float64:", v)
	case int:
		fmt.Println("int", v)
	}
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(8)
	DoSomething(12.0)
}
