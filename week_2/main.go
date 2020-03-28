package main

import (
	"fmt"
)

// AddI adds integers
type AddI func(int, int) int
type Gender int

// PI the PI you know
const PI = 3.14

const (
	Male Gender = iota
	Female
	Other
	NewGender
)

func main() {
	var tim Gender = 0
	var jane Gender = 1
	var they Gender = 2
	var NG Gender = 3
	var ug Gender = 4
	fmt.Printf("tim is a %s gender\n", tim.toString())
	fmt.Printf("jane is a %s gender\n", jane.toString())
	fmt.Printf("they is a %s gender\n", they.toString())
	fmt.Printf("NG is a %s gender\n", NG.toString())
	fmt.Printf("ug is a %s gender\n", ug.toString())
}

func (g Gender) toString() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	case Other:
		return "Other"
	case NewGender:
		return "NewGender"
	default:
		return "Unknown Gender"
	}
}

// gobyexample.com
