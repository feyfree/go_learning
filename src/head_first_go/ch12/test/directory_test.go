package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
		//panic(p)
	}
}

func scanDirectory(path string) {
	fmt.Println(path)
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range dir {
		filePath := filepath.Join(path, file.Name())
		if file.IsDir() {
			scanDirectory(filePath)
		} else {
			fmt.Println(filePath)
		}
	}
}

func TestDirectory(t *testing.T) {
	defer reportPanic()
	//panic("Some other issue")
	scanDirectory("GO")
}
