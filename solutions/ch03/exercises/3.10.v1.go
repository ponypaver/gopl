package main

import "fmt"

// 123456789 => 123,456,789
// 1234 => 1,234
func comma(s string) string {
	var ts string
	rs := reverse(s)

	for i:=0; i<len(rs); i+=3 {
		j := i+3
		sep := ","
		if j >= len(rs) {
			sep = ""
			j = len(rs)
		}
		ts += rs[i:j] + sep
	}

	return reverse(ts)
}

func reverse(src string) string {
	s := []rune(src)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(comma("123456789"))
}