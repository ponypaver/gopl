package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counts map[string]int
	files := os.Args[1:]

	if len(files) == 0 {
		counts = make(map[string]int)
		countLines(os.Stdin, counts)
		printDups(os.Stdin.Name(), counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			counts = make(map[string]int)
			countLines(f, counts)
			printDups(f.Name(), counts)
			f.Close()
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printDups(name string, counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%v: %s\t%d\n", name, line, count)
			//fmt.Printf("%d\t%s\n", count, line)
		}
	}
}