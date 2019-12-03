package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	c256 "crypto/sha256"
	c512 "crypto/sha512"
)

var (
	sha384 *bool
	sha512 *bool
)

func init() {
	sha384 = flag.Bool("sha384", false, "print sha384 hash")
	sha512 = flag.Bool("sha512", false, "print sha512 hash")
	flag.Parse()
}

func main() {
	//input := bufio.NewReader(os.Stdin)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		x := input.Text()

		switch {
		case *sha384:
			fmt.Printf("sha384(%v) = %x\n", x, c512.Sum384([]byte(x)))
		case *sha512:
			fmt.Printf("sha512(%v) = %x\n", x, c512.Sum512([]byte(x)))
		default:
			fmt.Printf("sha256(%v) = %x\n", x, c256.Sum256([]byte(x)))
		}
	}
}
