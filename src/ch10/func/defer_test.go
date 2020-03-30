package _func

import (
	"fmt"
	"testing"
)

func Clear() {
	fmt.Println("Clear Resources")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start")
	panic("Fatal Error")

}
