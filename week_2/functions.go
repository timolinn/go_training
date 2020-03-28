package main

import "fmt"

func functions() {
	r := addInts(20, 30)
	fmt.Printf("20 + 30 = %d\n", r)

	r = addInts(-5, 15)
	fmt.Printf("-5 + 15 = %d\n", r)

	var adder AddI

	adder = func(x, y int) int {
		return x + y
	}

	fmt.Printf("500 + 400 = %d\n", adder(500, 400))
}

func addInts(a, b int) int {
	// body
	return a + b
}
