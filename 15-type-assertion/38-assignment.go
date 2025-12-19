package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10))                                    // => 10
	fmt.Println(sum(10, "20"))                              // => 30
	fmt.Println(sum(10, 20, "30", 40, "abc"))               // => 100
	fmt.Println(sum([]int{10, 20}, []any{"30", 40, "abc"})) // => 100
	fmt.Println(sum())                                      // => 0
}

func sum(list ...any) int {
	var result int
	for _, item := range list {
		switch v := item.(type) {
		case int:
			result += v
		case string:
			if no, err := strconv.Atoi(v); err == nil {
				result += no
			}
		}
	}
	return result
}
