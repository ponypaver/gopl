package main

import (
	"fmt"
	"unicode"
)

// Write an in-place function that squashes each run of adjacent Unicode spaces
//(see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
func main() {
	s := "        走过      平湖         烟雨，   跨过    岁月    山河          "
	fmt.Println(mergeSpace(s))
}

func mergeSpace(s string) string {
	b := []byte(s)
	n := len(b)

	if n == 0 || n == 1 {
		return string(b)
	}

	tmp := b[:0]
	for i, j := 0, 1; j<n; i, j = i+1, j+1 {
		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1])) {
			continue
		}

		tmp = append(tmp, b[i])
	}

	return string(tmp)
}