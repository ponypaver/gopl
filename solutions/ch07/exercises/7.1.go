// Using the ide as from ByteCounter, implement counters for words and for lines.
// You will find bufio.ScanWords useful.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

type WordCounter int
type LineCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*c++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return int(*c), nil
}

//func (c *LineCounter) Write(p []byte) (int, error) {
//	scanner := bufio.NewScanner(bytes.NewReader(p))
//	scanner.Split(bufio.ScanLines)
//
//	for scanner.Scan() {
//		*c++
//	}
//
//	if err := scanner.Err(); err != nil {
//		return 0, err
//	}
//
//	return int(*c), nil
//}

func (c *LineCounter) Write(p []byte) (int, error) {

	*c = LineCounter(len(strings.Split(string(p), "\n")) - 1)

	return int(*c), nil
}

func main() {
	var w WordCounter
	_, _ = w.Write([]byte("hello world!"))
	fmt.Println(w)

	var l LineCounter
	_, _ = l.Write([]byte("hello world!\nabcd\ncd"))
	fmt.Println(l)
}