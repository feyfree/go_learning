package main

import "fmt"

func main() {
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}

	fmt.Printf("%#v\n", []rune("世界")) // []int32{19990, 30028}
	world := string([]rune{'世', '界'})
	for i, c := range world {
		fmt.Println(i, c)
	}
	fmt.Printf("%#v\n", world) // 世界
}
