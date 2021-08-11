package main

import "testing"

var a string
var done bool

func setup() {
	a = "hello, world"
	done = true
}

func Test(t *testing.T) {
	go setup()
	for !done {
	}
	print(a)
}

func TestChannelCsp(t *testing.T) {
	done := make(chan int)

	go func() {
		println("你好, 世界")
		done <- 1
	}()

	<-done
}
