package main

import "fmt"

func main() {
	hello := "hello World"
	l := hello[3:]

	if l == "lo World" {
		fmt.Println("True")
	}
	// fmt.Println()
}
