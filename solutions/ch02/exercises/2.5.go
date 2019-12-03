//The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
//of PopCount that counts bits by using this fact, and assess its per for \mance.

package main

import "fmt"

func popCount(x uint64) int {
	count := 0
	for ;x != 0; x = x&(x-1) {
		count++
	}
	return count
}

func main() {
	for i :=0; i< 10; i++ {
		fmt.Println(popCount(uint64(i)))
	}
}