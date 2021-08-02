package main

import (
	"fmt"
	"testing"
	"time"
)

func TestArrayDefinition(t *testing.T) {
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"你好", "世界"}
	var s3 = [...]string{1: "世界", 0: "你好"}
	iterAndPrintStringArray(s1[:])
	iterAndPrintStringArray(s2[:])
	iterAndPrintStringArray(s3[:])
}

func iterAndPrintStringArray(a []string) {
	fmt.Println("------start-------")
	// print type
	fmt.Printf("input: %T\n", a)
	// print value
	fmt.Printf("input: %#v\n", a)
	for i := range a {
		fmt.Printf("a[%d]: %s\n", i, a[i])
	}
	fmt.Printf("******end********")
}

func TestEmptyArray(t *testing.T) {
	// 长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，
	// 比如用于管道的同步操作
	c1 := make(chan [0]int)
	go func() {
		fmt.Println("c1")
		c1 <- [0]int{}
	}()
	<-c1

	// 在这里，我们并不关心管道中传输数据的真实类型，其中管道接收和发送操作只是用于消息的同步。
	// 对于这种场景，我们用空数组来作为管道类型可以减少管道元素赋值时的开销。当然一般更倾向于用无类型的匿名结构体代替
	c2 := make(chan struct{})
	go func() {
		fmt.Println("c2")
		time.Sleep(time.Second * 10)
		c2 <- struct{}{} // struct{}部分是类型, {}表示对应的结构体值
	}()
	<-c2
}
