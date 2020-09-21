package test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestWeb(t *testing.T) {
	response, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}

func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len(body))
}

func TestGoroutine(t *testing.T) {
	previous := time.Now().UnixNano()
	//go statements cannot be used with return values
	go responseSize("http://www.baidu.com")
	go responseSize("http://www.sohu.com")
	fmt.Println("spend:", time.Now().UnixNano()-previous)
	time.Sleep(time.Second)
}
