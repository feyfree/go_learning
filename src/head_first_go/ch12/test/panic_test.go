package test

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// 遇到处理不了的错误，找panic
//panic有操守，退出前会执行本goroutine的defer，方式是原路返回(reverse order)
//panic不多管，不是本goroutine的defer，不执行
func TestPanic(t *testing.T) {
	var user = os.Getenv("USER_")

	go func() {
		defer func() {
			fmt.Println("defer here")
		}()

		if user == "" {
			panic("should set user env.")
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Get Result")
}

func TestMultiDefer(t *testing.T) {
	defer fmt.Println("defer main") // will this be printed when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer fmt.Println("defer caller")
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Printf("get result")
}
