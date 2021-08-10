package main

import (
	"fmt"
	"sync"
)

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

type Cache struct {
	m map[string]string
	sync.Mutex
}

// Lookup Cache结构体类型通过嵌入一个匿名的sync.Mutex来继承它的Lock和Unlock方法.
// 但是在调用p.Lock()和p.Unlock()时, p并不是Lock和Unlock方法的真正接收者,
// 而是会将它们展开为p.Mutex.Lock()和p.Mutex.Unlock()调用. 这种展开是编译期完成的, 并没有运行时代价.
func (p *Cache) Lookup(key string) string {
	p.Lock()
	defer p.Unlock()

	return p.m[key]
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
