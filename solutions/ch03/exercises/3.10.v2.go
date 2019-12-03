package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	var i, j int
	n := len(s)
	buf := bytes.NewBufferString("")

	if n <= 3 {
		return s
	}

	for ; i < n; i = j {
		sep := ","
		j = i + 3
		if i == 0 {
			j = n % 3
			if j == 0 {
				j = 3
			}
		}

		if j == n {
			sep = ""
		}

		buf.WriteString(s[i:j])
		buf.WriteString(sep)
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234567889890909876"))
}
