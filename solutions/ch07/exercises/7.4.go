/*
The strings.NewReader function returns a value that satisfies the io.Reader
interface (and others) by reading from its argument, a string . Implement a simple version of
NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string
*/

package main

import (
	"fmt"
	"io"
)

type html struct {
	content string
	count   int
}

func (s *html) Read(p []byte) (n int, err error) {
	c := cap(p)
	if c == 0 {
		return 0, nil
	}

	if c >= len(s.content[s.count:]) {
		c = len(s.content[s.count:])
	}

	n = copy(p[:], s.content[s.count:c+s.count])
	s.count += c

	if s.count >= len(s.content) {
		return n, io.EOF
	}

	return n, nil
}

func NewHtml(s string) *html {
	return &html{content: s}
}

func NewReader(s string) io.Reader {
	return NewHtml(s)
}

func main() {
	r := NewReader("Hello, Reader!")

	b := make([]byte, 9)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
