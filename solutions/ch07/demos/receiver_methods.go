package main

import "fmt"

type T struct {
	name string
}

func (t *T) SetName(s string) {
	t.name = s
}
func (t T) bar() {
	t.name = "T.bar()"
}

func (t *T) String() string {
	return t.name
}

func main() {
	t := T{"init"}
	pt := &t

	// both T and *T could call a Pointer Receiver Method
	t.SetName("t init")
	// String() Method will not be called since T lacks String() Method
	fmt.Println(t)
	pt.SetName("pt init")
	// String() Method will not be called since T lacks String() Method
	fmt.Println(t)

	// reset to init
	pt.SetName("init")

	// both T and *T could call a Non-Pointer Receiver Method
	pt.bar()
	// String() Method called, *T has a String() Method
	fmt.Println(pt)
	pt.bar()
	// String() Method called, *T has a String() Method
	fmt.Println(pt)
}