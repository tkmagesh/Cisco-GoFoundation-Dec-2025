package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex
	count int
}

func (sc *SafeCounter) Add(delta int) {
	sc.Lock()
	{
		sc.count += 1
	}
	sc.Unlock()
}

var sf SafeCounter

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
	fmt.Println("Count :", sf.count)
}

func increment() {
	sf.Add(1)
}
