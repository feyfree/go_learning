package slice_test

import (
	"testing"
)

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	//t.Log(s2[4])
	s2 = append(s2, 1)
	s2 = append(s2, 2)
	t.Log(s2[3], s2[4])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(year)
}

func TestSliceCompare(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{1, 2, 3, 4}
	//slice 不可比较
	//if a == b {
	//	t.Log("Equal")
	//}
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	if a == b {
		t.Log("Equal")
	}
	t.Log(cap(a), len(a))
}

func TestSliceAndArray(t *testing.T) {
	arrayA := [2]int{100, 200}
	var arrayB [2]int
	arrayB = arrayA
	// 数组是值复制
	t.Log(&arrayB == &arrayA)
	// 不同类型的指针不能比较
	//sliceA := arrayA[:]
	//slice_addr := &sliceA
	//t.Log(slice_addr == &arrayA)
}
