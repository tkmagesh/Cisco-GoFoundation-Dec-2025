package main

import (
	"fmt"
	"sync"
)

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	primesCh := genPrimes(start, end)
	for primeNo := range primesCh {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func genPrimes(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		for no := start; no <= end; no++ {
			wg.Add(1)
			go checkPrime(no, wg, ch)
		}
		wg.Wait()
		close(ch)
	}()
	return ch
}

func checkPrime(no int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	ch <- no
}
