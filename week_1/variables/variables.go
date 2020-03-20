// Variable
package main

import "fmt"

var name = "Tom"

func main() {
	// Zero values int => 0, float => 0, string => "", bool => false
	// var number int = 100
	// j, k, l := 100, 120, 150
	var m, n int = 254, 12
	fmt.Printf("m = %d, in binary format = %b\n", m, m)
	fmt.Printf("n = %d, in binary format = %b\n", n, n)

	m, n = n, m

	fmt.Printf("m = %d, in binary format = %b\n", m, m)
	fmt.Printf("n = %d, in binary format = %b\n", n, n)
}
