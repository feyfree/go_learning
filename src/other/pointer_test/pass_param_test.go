package pointer

import (
	"fmt"
	"testing"
)

type Bird struct {
	Age  int
	Name string
}

type Parrot struct {
	Age  int
	Name string
}

// 传递T
func passV(b Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p\n", b, &b)
}

// 传递 *T
func passP(b *Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p, 指针的内存地址: %p\n", *b, b, &b)
}

var parrot1 = Bird{Age: 1, Name: "Blue"}
var parrot2 = parrot1

//////////////////////////////////////////////////////////////////////////////

func TestMemDiff(t *testing.T) {
	parrot := Bird{Age: 1, Name: "Blue"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
	passV(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
}

func TestMemDiffPassPointer(t *testing.T) {
	parrot := &Bird{Age: 1, Name: "Blue"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p, 指针的内存地址: %p\n", *parrot, parrot, &parrot)
	passP(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t内存地址：%p, 指针的内存地址: %p\n", *parrot, parrot, &parrot)
}

func TestBasicAssignment(t *testing.T) {
	fmt.Printf("parrot1:\t\t %+v, \t\t内存地址：%p\n", parrot1, &parrot1)
	fmt.Printf("parrot2:\t\t %+v, \t\t内存地址：%p\n", parrot2, &parrot2)
	parrot3 := parrot1
	fmt.Printf("parrot2:\t\t %+v, \t\t内存地址：%p\n", parrot3, &parrot3)
	parrot4 := Parrot(parrot1)
	fmt.Printf("parrot4:\t\t %+v, \t\t内存地址：%p\n", parrot4, &parrot4)
}

func TestMapSlice(t *testing.T) {
	fmt.Printf("parrot1:\t\t %+v, \t\t内存地址：%p\n", parrot1, &parrot1)
	//slice
	s := []Bird{parrot1}
	s = append(s, parrot1)
	parrot1.Age = 3
	fmt.Printf("parrot2:\t\t %+v, \t\t内存地址：%p\n", s[0], &(s[0]))
	fmt.Printf("parrot3:\t\t %+v, \t\t内存地址：%p\n", s[1], &(s[1]))
	parrot1.Age = 1
	//map
	m := make(map[int]Bird)
	m[0] = parrot1
	parrot1.Age = 4
	fmt.Printf("parrot4:\t\t %+v\n", m[0])
	parrot1.Age = 5
	parrot5 := m[0]
	fmt.Printf("parrot5:\t\t %+v, \t\t内存地址：%p\n", parrot5, &parrot5)
	parrot1.Age = 1
	//array
	a := [2]Bird{parrot1}
	parrot1.Age = 6
	fmt.Printf("parrot6:\t\t %+v, \t\t内存地址：%p\n", a[0], &a[0])
	parrot1.Age = 1
	a[1] = parrot1
	parrot1.Age = 7
	fmt.Printf("parrot7:\t\t %+v, \t\t内存地址：%p\n", a[1], &a[1])
}

func TestForRange(t *testing.T) {
	fmt.Printf("parrot1:\t\t %+v, \t\t内存地址：%p\n", parrot1, &parrot1)
	//slice
	s := []Bird{parrot1, parrot1, parrot1}
	s[0].Age = 1
	s[1].Age = 2
	s[2].Age = 3
	parrot1.Age = 4
	for i, p := range s {
		fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", i+2, p, &p)
	}
	parrot1.Age = 1
	//map
	m := make(map[int]Bird)
	parrot1.Age = 1
	m[0] = parrot1
	parrot1.Age = 2
	m[1] = parrot1
	parrot1.Age = 3
	m[2] = parrot1
	parrot1.Age = 4
	for k, v := range m {
		fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", k+2, v, &v)
	}
	parrot1.Age = 4
	//array
	a := [...]Bird{parrot1, parrot1, parrot1}
	a[0].Age = 1
	a[1].Age = 2
	a[2].Age = 3
	parrot1.Age = 4
	for i, p := range a {
		fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", i+2, p, &p)
	}
}

func TestChannel(t *testing.T) {
	ch := make(chan Bird, 3)
	fmt.Printf("parrot1:\t\t %+v, \t\t内存地址：%p\n", parrot1, &parrot1)
	ch <- parrot1
	parrot1.Age = 2
	ch <- parrot1
	parrot1.Age = 3
	ch <- parrot1
	parrot1.Age = 4
	p := <-ch
	fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", 2, p, &p)
	p = <-ch
	fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", 3, p, &p)
	p = <-ch
	fmt.Printf("parrot%d:\t\t %+v, \t\t内存地址：%p\n", 4, p, &p)
}

func TestArray(t *testing.T) {
	a1 := [3]int{1, 2, 3}
	fmt.Printf("a1:\t\t %+v, \t\t内存地址：%p\n", a1, &a1)
	a2 := a1
	a1[0] = 4
	a1[1] = 5
	a1[2] = 6
	fmt.Printf("a2:\t\t %+v, \t\t内存地址：%p\n", a2, &a2)
}

func TestFunc(t *testing.T) {
	f1 := func(i int) {}
	fmt.Printf("f1:  \t\t内存地址：%p\n", &f1)
	f2 := f1
	fmt.Printf("f2:  \t\t内存地址：%p\n", &f2)
}
