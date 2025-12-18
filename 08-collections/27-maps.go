package main

import "fmt"

func main() {
	// var productRanks map[string]int

	// var productRanks map[string]int = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// var productRanks = map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	// productRanks := map[string]int{"pen": 5, "pencil": 1, "marker": 3}
	/*
		productRanks := map[string]int{
			"pen":    5,
			"pencil": 1,
			"marker": 3,
		}
	*/

	var productRanks map[string]int
	productRanks = make(map[string]int)
	productRanks["pen"] = 5
	productRanks["pencil"] = 1
	productRanks["marker"] = 3
	fmt.Println(productRanks)

	// using key
	fmt.Printf("Rank of marker = %d\n", productRanks["marker"])

	// iterating using range
	for key, val := range productRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	// check for the existence of the key
	fmt.Println("check for the existence of the key")
	// keyToCheck := "pen"
	keyToCheck := "scribble-pad" // non-existent key

	// fmt.Printf("productRanks[%q] = %d\n", keyToCheck, productRanks[keyToCheck])
	if val, exists := productRanks[keyToCheck]; exists {
		fmt.Printf("productRanks[%q] = %d\n", keyToCheck, val)
	} else {
		fmt.Printf("key - %q does not exist\n", keyToCheck)
	}

	// remove a key ( using delete() )
	// keyToRemove := "pen"
	keyToRemove := "scribble-pad" // non-existent key, delete() is a no-op
	delete(productRanks, keyToRemove)
	fmt.Printf("productRanks after removing %q = %v\n", keyToRemove, productRanks)

	// like slices, maps are also pointers to underlying data
	newProductRanks := productRanks
	newProductRanks["pen"] = 10
	fmt.Printf("newProductRanks[pen] = %d, productRanks[pen] = %d\n", newProductRanks["pen"], productRanks["pen"])

}
