package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) []byte {

	if len(b) == 1 || len(b) == 0 {
		return b
	}

	_, size := utf8.DecodeRune(b)
	return append(reverse(b[size:]), b[:size]...)
}

func main() {
	fmt.Println(string(reverse([]byte("hello, 世界！"))))
}