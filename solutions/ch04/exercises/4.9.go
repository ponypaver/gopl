//Wr ite a program wordfreq to rep ort the frequency of each word in an inp ut text
//file. Cal l input.Split(bufio.ScanWords) before the firs t call to Scan to bre ak the inp ut int o
//word s instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	file, err := os.Open("words")
	if err != nil {
		file.Close()
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counter[input.Text()]++
	}
	if err := input.Err(); err != nil {
		file.Close()
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}

	file.Close()

	fmt.Printf("%20v\tcount\n", "word")
	for w, c := range counter {
		fmt.Printf("%20v\t%v\n", w, c)
	}
}
