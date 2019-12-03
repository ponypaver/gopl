package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go spinner(100 * time.Millisecond, ch)
	const n = 45
	fibN := fib(n) // slow
	ch <- true
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration, ch <- chan bool) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
		select {
		case <-ch:
			fmt.Println("\rFibonacci computation done")
			return
		default:
			continue
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
