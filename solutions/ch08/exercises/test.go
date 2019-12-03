package main

import (
	"fmt"
	"sync"
)

func main() {
	go have()(fun("with Go."))

	fmt.Print("some ") // evaluation order: ~ 3
	wg.Wait()
}

func have() func(string) {
	fmt.Print("Go ") // evaluation order: 1
	return funWithGo
}

func fun(msg string) string {
	fmt.Print("have ") // evaluation order: 2
	return msg
}

func funWithGo(msg string) {
	fmt.Println("fun", msg) // evaluation order: 4
	wg.Done()
}

func init() {
	wg.Add(1)
}

var wg sync.WaitGroup