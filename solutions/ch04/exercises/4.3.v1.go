package main

import "fmt"

// reverse reverses a slice of ints in place.
// func reverse(s []int) {
//	 for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
//	 	s[i], s[j] = s[j], s[i]
//	 }
// }

func reverse(p *[10]int) (re [10]int) {
	for i, j := len(p)-1, 0; i >= 0; i, j = i-1, j+1 {
		re[j] = p[i]
	}

	return
}

func main() {
	fmt.Println(reverse(&[10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}))
}
