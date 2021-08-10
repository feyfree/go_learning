package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

var totalValue uint64

func myWorker(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&totalValue, i)
	}
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go myWorker(&wg)
	go myWorker(&wg)
	wg.Wait()

	t.Log(totalValue)
}
