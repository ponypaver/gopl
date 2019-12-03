package main

import (
	"crypto/sha256"
	"fmt"
)

//Write a function that counts the number of bits that are different in two SHA256
//hashes. (See PopCount from Section 2.6.2.)

func countDiffBit(a, b byte) int {
	if a == b {
		return 0
	}

	count := 0
	for a != 0 || b != 0 {
		if a & 1 != b & 1 {
			count++
		}
		a = a >> 1
		b = b >> 1
	}

	return count
}

func main() {
	count := 0
	x := sha256.Sum256([]byte("x"))
	y := sha256.Sum256([]byte("y"))

	for i := range x {
		count += countDiffBit(x[i], y[i])
	}

	fmt.Println(count)

}