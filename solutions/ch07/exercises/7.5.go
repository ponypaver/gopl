/*
The LimitReader function in the io package accepts an io.Reader r and a
number of bytes n, and returns another Reader that reads from r but reports an end-of-file
condition after n bytes. Implement it.
*/

package main

import (
	"fmt"
	"io"
	//"os"
	//"log"
	"strings"
)

type limitedReader struct {
	r     io.Reader
	count int
	n     int
}

func newLimitedReader(r io.Reader, n int) *limitedReader {
	return &limitedReader{r, 0, n}
}

func (r *limitedReader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	if r.count + n >= r.n {
		return r.n - r.count, io.EOF
	}

	r.count += n
	return n, err
}

func limitReader(r io.Reader, n int) io.Reader {
	return newLimitedReader(r, n)
}

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	//lr := io.LimitReader(r, 4)
	lr := limitReader(r, 9)

	//if _, err := io.Copy(os.Stdout, lr); err != nil {
	//	log.Fatal(err)
	//}

	b := make([]byte, 1)
	for {
		n, err := lr.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
