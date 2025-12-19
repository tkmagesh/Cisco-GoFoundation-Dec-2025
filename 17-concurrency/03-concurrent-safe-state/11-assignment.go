package main

import (
	"fmt"
	"sync"
)

/*
Modify as below

1. Do not print the prime number in "checkPrime()"
2. Collect all the prime numbers in a collection
3. After all the prime numbers are identified, print them one after another in the main()
*/

// communicate by sharing memory
var primes []int

var mutex sync.Mutex

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
	wg := &sync.WaitGroup{}
	for no := start; no <= end; no++ {
		wg.Add(1)
		go checkPrime(no, wg)
	}
	wg.Wait()
	for _, primeNo := range primes {
		fmt.Println("Prime No :", primeNo)
	}
	fmt.Println("Done")
}

func checkPrime(no int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= (no / 2); i++ {
		if no%i == 0 {
			return
		}
	}
	mutex.Lock()
	{
		primes = append(primes, no)
	}
	mutex.Unlock()
}
