package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func main() {
	var wg sync.WaitGroup
	for range 300 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("Count :", count)
}

func increment() {
	atomic.AddInt64(&count, 1)
}
