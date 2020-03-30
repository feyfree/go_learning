package string_test

import "testing"

/*
string 是不可变的byte slice
*/
func TestString(t *testing.T) {
	var s string
	s = "hello"
	t.Log(len(s))
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "\xE4\xBB\xA5"
	t.Log(s)
	s = "中"
	t.Log(len(s))
	c := []rune(s)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 utf8 %x", s)
}

func TestStringRune(t *testing.T) {
	s := "中华人民共和国"
	for _, v := range s {
		t.Logf("%[1]c %[1]x", v)
	}
}
