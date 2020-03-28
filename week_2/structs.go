package main

import (
	"fmt"
)

// COV19Stats represents the COVID19 status
type COV19Stats struct {
	Country       string
	NumberOfCases int
	Safety        SafetyMeasures
}

// SafetyMeasures represents the necessary precauations
// to reduce the spread of COVID19
type SafetyMeasures struct {
	First  string
	Second string
	Third  string
}

func structs() {
	// nigeria := COV19Stats{
	// 	Country:       "Nigeria",
	// 	NumberOfCases: 81,
	// 	Safety:        SafetyMeasures{},
	// }
	nigeria := COV19Stats{}
	nigeria.Country = "Nigeria"
	nigeria.NumberOfCases = 81
	fmt.Println(nigeria.Country)
	SetCountry(&nigeria, "Ghana")
	fmt.Println(nigeria.Country)
}

// SetCountry sets a new country value
func SetCountry(cnt *COV19Stats, value string) {
	cnt.Country = value
}

// gobyexamples.com
