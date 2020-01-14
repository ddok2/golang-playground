package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fib(n)

	fmt.Printf("\rFibonacci (%d) = %d", n, fibN)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
