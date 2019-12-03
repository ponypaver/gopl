package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)

	go func() {
		time.Sleep(time.Second*5)
		close(ch)
	}()

	select {
	case x := <-ch:
		fmt.Printf("received on ch: %v\n", x)
	}
}
