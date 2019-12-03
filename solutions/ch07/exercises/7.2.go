// Write a function CountingWriter with the signature below that, given an
// io.Writer, returns a new Writer that wraps the original, and a pointer to an int64 variable
// that at any moment contains the number of bytes written to the new Writer.
//
// func CountingWriter(w io.Writer) (io.Writer, *int64)

package main

import (
	"fmt"
	"io"
	"io/ioutil"
)

type ByteCounter struct {
	w     io.Writer
	count int64
}

func (c *ByteCounter) Write(p []byte) (n int, err error) {
	n, err = c.w.Write(p)
	c.count += int64(n)

	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	bc := ByteCounter{w: w}

	return &bc, &bc.count
}

func main() {
	w, c := CountingWriter(ioutil.Discard)

	w.Write([]byte("滚滚红尘，情缘如梦，如花凋零\n"))
	w.Write([]byte("无奈流水匆匆\n"))

	fmt.Println(*c)
}
