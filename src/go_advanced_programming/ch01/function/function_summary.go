package main

import "fmt"

func Add(a, b int) int {
	var add = func(a, b int) int {
		return a + b
	}
	result := add(a, b)
	fmt.Println(result)
	return result
}

func Swap(a, b int) (int, int) {
	return b, a
}

// Sum 可变数量的参数 more 对应 []int 切片类型
func Sum(a int, more ...int) int {
	for _, v := range more {
		a += v
	}
	fmt.Println(a)
	return a
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {
	a := 1
	b := 2
	Add(a, b)
	Swap(a, b)
	Sum(a, a, b)

	var c = []interface{}{123, "abc"}
	// 第一个Print调用时传入的参数是a...，等价于直接调用Print(123, "abc")。
	// 第二个Print调用传入的是未解包的a，等价于直接调用Print([]interface{}{123, "abc"})。
	Print(c...) // 123 abc
	Print(c)    // [123 abc]

}
