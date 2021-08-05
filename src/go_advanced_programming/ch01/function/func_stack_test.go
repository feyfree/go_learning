package main

import "testing"

func f(x int) *int {
	return &x
}

func g() int {
	x := new(int)
	return *x
}

func TestFuncStack(t *testing.T) {

}
