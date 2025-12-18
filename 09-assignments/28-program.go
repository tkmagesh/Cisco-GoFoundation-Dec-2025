package assignments
/*
Refactor as follows:

	IsPrime(no) bool 
		- return true if the given number is a prime number
		
	GenPrimes(start, end) ?
		- return a collection of prime numbers between 'start' and 'end'
	
	main()
		receive the prime numbers collection from GenPrimes() and print them one after another


*/

package main

import "fmt"

func main() {
	var start, end int
	fmt.Println("Enter the start and end :")
	fmt.Scanln(&start, &end)
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i <= (no / 2); i++ {
			if no%i == 0 {
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
