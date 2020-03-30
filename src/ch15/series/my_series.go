package series

import "fmt"

func init() {
	fmt.Println("init1")
}

func init() {
	fmt.Println("init2")
}

// 小写的函数声明是不能被包外访问的
func Square(n int) int {
	return n * n
}

func GetFibonacciSeries(n int) []int {
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
