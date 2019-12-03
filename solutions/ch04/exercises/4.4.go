package main

import "fmt"

// rotate rotate a slice left by n elements
// s = [0 1 2 3 4 5]
// n = 2
// -> [2 3 4 5 0 1]
func rotate(s []int, n int) {
	l := len(s)
	//ts := make([]int, l)

	for i := 0; i < n%l; i++ {
		tmp := s[0]
		copy(s[0:l-1], s[1:])
		s[l-1] = tmp
	}
	//fmt.Println(ts)
	//s = ts
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	rotate(s, 7)
	fmt.Println(s)
}
