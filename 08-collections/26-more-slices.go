package main

import "fmt"

func main() {
	// var nos []int

	// preallocate using make()
	// nos := make([]int, 0, 5)

	nos := make([]int, 3, 5)

	// capacity = memory allocated to the underlying array
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 3)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 5)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 2)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 1)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 4)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)

	nos = append(nos, 7)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %d\n", len(nos), cap(nos), nos)
}
