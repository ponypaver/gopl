/*
Use panic and recover to write a function that contains no return statement
yet returns a non-zero value.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(foo())
}

func foo() (ret int) {
	defer func() {
		p := recover()
		ret, _ = strconv.Atoi(fmt.Sprintf("%v", p))
	}()

	panic(1)
}
