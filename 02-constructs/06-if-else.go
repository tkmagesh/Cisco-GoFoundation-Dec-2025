package main

import "fmt"

func main() {
	/*
		no := 21
		if no%2 == 0 {
			fmt.Printf("%d is an even number\n", no)
		} else {
			fmt.Printf("%d is an odd number\n", no)
		}
	*/

	if no := 21; no%2 == 0 { // "no" variable scope is limited to "if-else" block
		fmt.Printf("%d is an even number\n", no)
	} else {
		fmt.Printf("%d is an odd number\n", no)
	}

	// "no" variable is NOT accessible outside the "if-else" block
	// fmt.Println("no = ", no)
}
