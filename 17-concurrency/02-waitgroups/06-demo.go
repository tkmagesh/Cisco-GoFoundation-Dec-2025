package main

import (
	"fmt"
	"sync"
)

/*
Make  the logic of checking each number prime execute concurrently
*/

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if IsPrime(no) {
				fmt.Println("Prime :", no)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Done")
}

func IsPrime(no int) bool {
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
