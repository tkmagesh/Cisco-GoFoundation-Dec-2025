package main

import "fmt"

func main() {
	// declaration
	// var nos []int

	// initializing with our own data
	// var nos []int = []int{3, 1, 5, 2, 4}

	// type inference
	// var nos = []int{3, 1, 5, 2, 4}

	// using :=
	nos := []int{3, 1, 5, 2, 4}
	fmt.Println(nos)

	// using append()
	nos = append(nos, 40)
	nos = append(nos, 10, 30, 20)
	hundreds := []int{500, 200, 400}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	// using len()
	fmt.Printf("len(nos) = %d\n", len(nos))

	// iterating using indexer
	fmt.Println("iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	// iterating using range
	fmt.Println("iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	//
	newNos := nos
	newNos[0] = 9999
	fmt.Printf("newNos[0] = %d, nos[0] = %d\n", newNos[0], nos[0])

	fmt.Printf("Before sorting, nos = %v\n", nos)
	bubbleSort(nos)
	fmt.Printf("After sorting, nos = %v\n", nos)

	// slicing
	fmt.Printf("nos[3:6] = %v\n", nos[3:6])
	fmt.Printf("nos[:6] = %v\n", nos[:6])
	fmt.Printf("nos[3:] = %v\n", nos[3:])
}

func bubbleSort(list []int) { // no return values
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
