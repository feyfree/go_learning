package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针

	fmt.Println(a[0], a[1]) // 打印数组的前2个元素
	fmt.Println(b[0], b[1]) // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range b { // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	for i := range a {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]: %d\n", i, v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d]: %d\n", i, a[i])
	}

	var times [5][0]int
	for range times {
		fmt.Println("hello")
	}
}
