package main

// Package solution converts multiple units
// bi-directionally. eg. Celsius to Fahrenheit and
// Fahrenheit to Celsius

import "math"

type Fahrenheit float64
type Celsius float64
type Radian float64
type Degree float64

type Converter struct {
}

// FahrenheitToCelsius converts a Fahrenheit to Celsius
func (cvr Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// CelsiusToFahrenheit converts celsius to fahrenheit
func (cvr Converter) CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// RadianToDegree converts radian to degree
func (cvr Converter) RadianToDegree(r Radian) Degree {
	return Degree(r * (180 / math.Pi))
}

// DegreeToRadian converts degree to radian
func (cvr Converter) DegreeToRadian(d Degree) Radian {
	return Radian(d * (math.Pi / 180))
}
