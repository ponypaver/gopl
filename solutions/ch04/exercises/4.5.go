package main

import "fmt"

// Write an in-place function to eliminate adjacent duplicates in a []string slice
func eliAdjDup(s []string) []string {
	n, originLen := len(s), len(s)
	removed := 0

	if n == 0 || n == 1 {
		return s
	}

	for i, j := 0, 1; j < n; {
		if s[i] == s[j] {
			copy(s[i:], s[j:])
			removed++
			n--
			if removed == n-1 {
				break
			}
			continue
		}

		i, j = i+1, j+1
	}

	return s[:originLen-removed]
}

func main() {
	//s := []string{"hello", "a", "abcd", "abcd", "abc", "efg", "efg", "abcd", "a", "a", "b"}
	//s := []string{"a", "a", "a", "a"}
	s := []string{"x", "a", "a", "a", "b", "c", "c"}
	re := eliAdjDup(s)
	fmt.Println(re)
}
