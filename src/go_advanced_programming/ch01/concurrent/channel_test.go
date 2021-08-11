package main

import (
	"fmt"
	"testing"
)

var end = make(chan bool)
var msg string

func aGoroutine() {
	msg = "你好, 世界"
	close(end)
}

func TestChannel(t *testing.T) {
	go aGoroutine()
	<-end
	println(msg)
}

var limit = make(chan int, 3)

var work = make([]int, 3)

func do(a int) { fmt.Println("working now!") }

func TestLimitGoroutine(t *testing.T) {
	for _, w := range work {
		w := w
		go func() {
			limit <- 1
			do(w)
			<-limit
		}()
	}
	select {}
}
