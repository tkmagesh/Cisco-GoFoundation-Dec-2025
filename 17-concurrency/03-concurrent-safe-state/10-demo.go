package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count atomic.Int64

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
	fmt.Println("Count :", count.Load())
}

func increment() {
	count.Add(1)
}
