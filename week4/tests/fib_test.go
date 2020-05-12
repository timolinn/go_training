package tests

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	want := 13
	t.Run("fibonacci of 7 should be 13", func(t *testing.T) {
		got := Fibonacci(7)
		if got != want {
			t.Fatalf("wanted %d got %d", want, got)
		}
	})

	t.Run("fibonacci of 5 should be 5", func(t *testing.T) {
		got := Fibonacci(5)
		if got != 5 {
			t.Fatalf("wanted %d got %d", 5, got)
		}
	})
}
