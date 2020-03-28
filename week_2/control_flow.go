package main

import "fmt"

func controlFlow() {
	// if a == b // true or false {
	// 	fmt.Println("a is equal to b")
	// } else if a > b {
	// 	fmt.Println("a is greater than b")
	// } else if b < a {
	// 	fmt.Println("b is less than a")
	// } else {
	// 	fmt.Println("b is greater than a")
	// }

	switch add(4, 5) {
	case 8:
		fmt.Println("the result is 8")
		break
	case 9:
		fmt.Println("the result is 9")
		break
	case 10:
		fmt.Println("the result is 10")
		break
	default:
		fmt.Println("we don't know the answer to that")
	}
}

func add(a, b int) int {
	return a + b
}
