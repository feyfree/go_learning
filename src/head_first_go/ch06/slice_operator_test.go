// slice 相关操作
package ch06

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceCut(t *testing.T) {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	// [start, end] start 包含 end 不包含
	slice := underlyingArray[0:3]
	fmt.Println(slice)
	fmt.Println(reflect.TypeOf(slice))

	slice2 := underlyingArray[0:5]
	fmt.Println(slice2)

	slice3 := underlyingArray[1:]
	fmt.Println(slice3)

	underlyingArray[2] = "o"
	fmt.Println(slice2)
}

// slice 对应一个 array  append 可能会有新的 underlying array
func TestSliceAppend(t *testing.T) {
	s1 := []string{"s1", "s1"}
	s2 := append(s1, "s2", "s2")
	s3 := append(s2, "s3", "s3")
	s4 := append(s3, "s4", "s4")
	fmt.Println(s1, s2, s3, s4)
	s4[0] = "XX"
	fmt.Println(s1, s2, s3, s4)
}
