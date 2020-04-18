package week3

import (
	"fmt"
	"runtime"
	"time"
)

func goroutines() {
	runtime.GOMAXPROCS(2)
	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			println("Hello")
			time.Sleep(godur)

		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)

	// go spinner(100 * time.Millisecond)
	// const n = 45
	// fibN := fibonacci(n) // slow
	// fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
