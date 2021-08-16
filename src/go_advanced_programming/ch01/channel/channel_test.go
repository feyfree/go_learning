package channel

import (
	"fmt"
	"testing"
)

// 声明一个只能读取数据的通道类型
func printVal(a <-chan int) {
	fmt.Println(a)
	<-a
}

func printVal2(a chan int) {
	fmt.Println(<-a)
}

func TestChannelParam(t *testing.T) {
	a := make(chan int)
	a <- 1
	printVal2(a)
	a <- 1
	printVal(a)
}
