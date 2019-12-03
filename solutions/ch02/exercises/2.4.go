//Write a version of PopCount that counts bits by shifting its argument through 64
//bit positions, testing the rightmost bit each time. Compare its performance to the table lookup version.

package main

import "fmt"

func popCount(x uint64) int {
	count := 0
	for ; x != 0; x >>= 1 {
		if x&1 == 1 {
			count++
		}
	}

	return count
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(popCount(uint64(i)))
	}
}
