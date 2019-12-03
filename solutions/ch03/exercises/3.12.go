package main

import (
	"fmt"
	"strings"
)

func anagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, c := range a {
		if !strings.ContainsRune(b, c) {
			return false
		}

		b = strings.Replace(b, string(c), "", 1)
	}

	return true
}

func main() {
	fmt.Println(anagrams("aabbc", "abccc"))
}
