package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	split := strings.Split(s, ",")
	for _, part := range split {
		t.Log(part)
	}
	t.Log(strings.Join(split, "-"))
}

func TestStringConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str:", s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
