package main

import "fmt"

func main() {

	/*
		const red int = 0
		const green int = 1
		const blue int = 2
	*/

	/*
		const (
			red   int = 0
			green int = 1
			blue  int = 2
		)
	*/

	/*
		const (
			red   int = iota
			green int = iota
			blue  int = iota
		)
	*/

	/*
		const (
			red int = iota
			green
			blue
		)
	*/

	// type inference
	/*
		const (
			red = iota
			green
			blue
		)
	*/

	// skip
	/*
		const (
			red = iota
			green
			_
			blue
		)
	*/

	const (
		red   = (iota + 1) * 2 // (0+1)*2
		green                  // (1+1)*2
		_                      // (2+1)*2
		blue                   // (3+1)*2
	)
	fmt.Printf("red = %d, green = %d & blue = %d\n", red, green, blue)

	// Mimicking unix file permissions
	const (
		X = 1 << iota
		W
		R

		XW  = X | W
		WR  = W | R
		XWR = X | W | R
	)
	fmt.Printf("%b %b %b %b %b %b\n", X, W, XW, R, WR, XWR)
}
