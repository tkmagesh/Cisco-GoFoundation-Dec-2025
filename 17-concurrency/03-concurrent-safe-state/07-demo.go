package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

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
	mutex.Lock()
	{
		count += 1
	}
	mutex.Unlock()
}
