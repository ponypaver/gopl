package main

import (
	c256 "crypto/sha256"
	c512 "crypto/sha512"
	"flag"
	"fmt"
	"os"
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
	for _, x := range os.Args[1:] {
		switch {
		case *sha384:
			if x == "-sha384" {
				continue
			}
			fmt.Printf("sha384(%v) = %x\n", x, c512.Sum384([]byte(x)))
		case *sha512:
			if x == "-sha512" {
				continue
			}
			fmt.Printf("sha512(%v) = %x\n", x, c512.Sum512([]byte(x)))
		default:
			fmt.Printf("sha256(%v) = %x\n", x, c256.Sum256([]byte(x)))
		}
	}
}
