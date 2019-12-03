package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) error {
	defer c.Close()
	for {
		//if _, err := c.Write([]byte(time.Now().Format("13:01:01"))); err != nil {
		if _, err := io.WriteString(c, time.Now().Format("15:04:05\n")); err != nil {
			return err
		}

		time.Sleep(time.Second)
	}
}

func main() {
	listen, err := net.Listen("tcp","localhost:8000")

	if err != nil {
		log.Fatal()
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}