package test

import (
	"fmt"
	"testing"
)

func greeting(myChannel chan string) {
	myChannel <- "Hi!"
}

func TestSampleChannel(t *testing.T) {
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)

}
