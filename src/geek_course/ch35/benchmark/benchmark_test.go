package benchmark_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcatStringByAdd(t *testing.T) {
	assertion := assert.New(t)
	elms := []string{"1", "2", "3", "4", "5"}
	ret := ""
	for _, elem := range elms {
		ret += elem
	}
	assertion.Equal("12345", ret)
}

// 测试命令
// go test -bench=. -benchmem
func TestConcatStringByBytesBuffer(t *testing.T) {
	assertion := assert.New(t)
	var buf bytes.Buffer
	elms := []string{"1", "2", "3", "4", "5"}
	for _, elem := range elms {
		buf.WriteString(elem)

	}
	assertion.Equal("12345", buf.String())
}

func BenchmarkConcatStringByAdd(b *testing.B) {

	elms := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elms {
			ret += elem
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()

}
