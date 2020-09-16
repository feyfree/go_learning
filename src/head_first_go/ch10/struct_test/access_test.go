package struct_test

import (
	"fmt"
	"testing"
)

// go 语言对大小写敏感
type MyDate struct {
	year  int
	month int
	day   int
}

type EmbeddedType string

func (e *EmbeddedType) ExportMethod() {
	fmt.Println("Hi, This is an exported method !")
}

func (e *EmbeddedType) exportMethod() {
	fmt.Println("Hi, This is an exported method !")
}

type MyType struct {
	EmbeddedType
}

func TestMyType(t *testing.T) {
	myType := MyType{}
	myType.ExportMethod()
	myType.exportMethod()
}

func NewMyDate(year int, month int, day int) *MyDate {
	return &MyDate{year: year, month: month, day: day}
}

func TestMyDate(t *testing.T) {
	date := MyDate{2020, 9, 16}
	fmt.Println(date)
	newDate := NewMyDate(2020, 9, 16)
	fmt.Println(*newDate)
	fmt.Println(date.day)
}
