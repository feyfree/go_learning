package mytype

import "fmt"

type MyEmbeddedType string

func (t *MyEmbeddedType) ExportMethod() {
	fmt.Println("Hi, This is an exported method upper case!")
}

func (t *MyEmbeddedType) exportMethod() {
	fmt.Println("Hi, This is an exported method lower case!")
}

type MyNewType struct {
	MyEmbeddedType
}
