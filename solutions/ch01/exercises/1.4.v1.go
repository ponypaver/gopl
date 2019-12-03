package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	var inputs []*os.File
	files := os.Args[1:]

	if len(files) == 0 {
		inputs = []*os.File{os.Stdin}
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			inputs = append(inputs, f)
		}
	}

	for _, f := range inputs {
		countLines(f, counts)
		printDups(f, counts)
		_ = f.Close()
	}
}

func countLines(f *os.File, counter map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counter[input.Text()]++
	}
}

func printDups(f *os.File, counts map[string]int) {
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%v: %s\t%d\n", f.Name(), line, count)
			//fmt.Printf("%d\t%s\n", count, line)
		}
	}
}