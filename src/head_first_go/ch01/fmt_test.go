package ch01

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf(3.1415))
	fmt.Println(reflect.TypeOf(true))
}
