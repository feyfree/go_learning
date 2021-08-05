package main

import (
	"fmt"
	"testing"
)

func EnclosureA() {
	// Loop variables captured by 'func' literals in 'defer' statements might have unexpected values
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}

}
func EnclosureB() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() { println(i) }()
	}

}
func EnclosureC() {
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}

func TestEnclosure(t *testing.T) {
	EnclosureA()
	fmt.Println("----This is separation--------")
	EnclosureB()
	fmt.Println("----This is separation--------")
	EnclosureC()

}
