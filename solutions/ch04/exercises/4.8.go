//Mo dif y charcount to count letters, dig its, and so on in their Unico de categories,
//using functions like unicode.IsLetter.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letters := make(map[rune]int)
	digits := make(map[rune]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			fmt.Fprintf(os.Stderr, "charcount: invalide char: %v\n", r)
			continue
		}
		if unicode.IsLetter(r) {
			letters[r]++
		}
		if unicode.IsDigit(r) {
			digits[r]++
		}
	}
	fmt.Printf("letter\tcount\n")
	for l, c := range letters {
		fmt.Printf("%c\t%v\n", l,c )
	}

	fmt.Printf("digits\tcount\n")
	for l, c := range digits {
		fmt.Printf("%c\t%v\n", l,c )
	}
}