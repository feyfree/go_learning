package main

import "fmt"

func main() {
	//arrayA := [2]int{100, 200}
	//var arrayB [2]int
	//
	//arrayB = arrayA
	//
	//fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
	//fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)
	//
	//testArray(arrayA)
	arrayA := []int{100, 200}
	testArrayPoint(&arrayA) // 1.传数组指针
	arrayB := arrayA[:]
	testArrayPoint(&arrayB) // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArray(x [2]int) {
	fmt.Printf("func Array : %p , %v\n", &x, x)
}

func testArrayPoint(x *[]int) {
	fmt.Printf("func Array : %p , %v\n", x, *x)
	(*x)[1] += 100
}
