package main

import (
	"github.com/prometheus/common/log"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src net.Conn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
