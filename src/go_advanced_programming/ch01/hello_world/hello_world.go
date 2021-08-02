package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	printHelloWorld()
	helloWorldFromNet()
}

func printHelloWorld() {
	fmt.Println("Hello,World!")
}

func helloWorldFromNet() {
	fmt.Println("Please visit http://127.0.0.1:12345/")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("你好, 世界! -- Time: %s", time.Now().String())
		_, err := fmt.Fprintf(w, "%v\n", s)
		if err != nil {
			return
		}
		log.Printf("%v\n", s)
	})
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
