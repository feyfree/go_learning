package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	return p.Writer.Write(bytes.ToUpper(data))
}

type TB struct {
	// embedded type cannot be pointer to an interface
	testing.TB
}

func (p *TB) Fatal(args ...interface{}) {
	fmt.Println("TB.Fatal disabled!")
}

func main() {
	_, err := fmt.Fprintln(&UpperWriter{os.Stdout}, "hello, world")
	if err != nil {
		return
	}

	var tb testing.TB = new(TB)
	tb.Fatal("Hello, playground")
}
