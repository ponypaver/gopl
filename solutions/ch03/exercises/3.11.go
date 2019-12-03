package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := os.Args[1:]

	for j := range numbers {
		_, err := strconv.ParseInt(numbers[j], 10, 64)
		if err != nil {
			_, err := strconv.ParseFloat(numbers[j], 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s => invalid number: %v\n", numbers[j], numbers[j])
				continue
			}
		}

		var re []string
		for _, part := range strings.Split(numbers[j], ".") {
			re = append(re, comma(part))
		}

		fmt.Printf("%s => %s\n", numbers[j], strings.Join(re, "."))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
