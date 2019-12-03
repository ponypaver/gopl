package main

import (
	"fmt"
	"os"
	"sync"
)

var cancel = make(chan struct{})

func cancelled() bool {
	select {
	case <-cancel:
		return true
	default:
		return false
	}
}

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	go func() {
		defer wg.Done()
		for {
			if cancelled() {
				fmt.Println("cancelled")
				break
			}
		}
	}()

	wg.Wait()

}
