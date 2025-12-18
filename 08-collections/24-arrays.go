package main

import "fmt"

func main() {
	//memory allocated and initialized with zero value of 'int'
	// var nos [5]int

	// initializing with our own data
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}

	// type inference
	// var nos = [5]int{3, 1, 4, 2, 5}

	// using :=
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	// accessing the data using indexers
	fmt.Println("using indexers")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// accessing the data using 'range'
	fmt.Println("using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	// assignment will result in a copy being created coz "everything is a value in go"
	newNos := nos
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d, nos[0] = %d\n", newNos[0], nos[0])

	// handling references of an array
	var noPtr *[5]int
	noPtr = &nos

	// accessing the members from a "pointer to an array" (dereferencing)
	fmt.Printf("(*noPtr)[0] = %d\n", (*noPtr)[0])

	// Accessing the members of a "pointer to a container" DOES NOT require dererencing
	fmt.Printf("noPtr[0] = %d\n", noPtr[0])

	fmt.Printf("Before sorting, nos = %v\n", nos)
	bubbleSort(&nos)
	fmt.Printf("After sorting, nos = %v\n", nos)
}

func bubbleSort(list *[5]int) { // no return values
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
