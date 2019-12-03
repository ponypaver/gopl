// Rewrite PopCount to use a loop instead of a single expression.
// Compare the performance of the two versions.
// (Section 11.4 shows how to compare the performance of different implementations systematically.)

package main

import "fmt"

// pc[i] is the population count of i.
//var pc [256]byte
//func init() {
//	for i := range pc {
//		pc[i] = pc[i/2] + byte(i&1)
//	}
//}
//// PopCount returns the population count (number of set bits) of x.
//func PopCount(x uint64) int {
//	return int(pc[byte(x>>(0*8))] +
//		pc[byte(x>>(1*8))] +
//		pc[byte(x>>(2*8))] +
//		pc[byte(x>>(3*8))] +
//		pc[byte(x>>(4*8))] +
//		pc[byte(x>>(5*8))] +
//		pc[byte(x>>(6*8))] +
//		pc[byte(x>>(7*8))])
//}

var pc [256]byte
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func popCount(x uint64) int {
	count := 0
	for ; x != 0; x >>= 8 {
		count += int(pc[byte(x)])
	}
	return count
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(popCount(uint64(i)))
	}
}