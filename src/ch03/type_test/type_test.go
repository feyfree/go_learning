package type_test

import "testing"

type MyInt int64

/**
go 不支持隐式 类型转换
*/
func TestImplicit(t *testing.T) {
	//var a int = 1
	//var b int64;
	//b = a
	//t.Log(a, b)

	//var a int32 = 1
	//var b int64;
	//b = a
	//t.Log(a, b)

	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// Go 语言不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

}
