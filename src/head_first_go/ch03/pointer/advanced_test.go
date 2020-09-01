package pointer

import (
	"fmt"
	"testing"
)

func doubleByPointer(number *int) {
	*number = *number * 2
}

func TestDouble(t *testing.T) {
	amount := 6
	doubleByPointer(&amount)
	fmt.Println(amount)
}
