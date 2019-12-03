/*
Write a String method for the *tree type in gopl.io/ch4/treesort (§4.4)
that reveals the sequence of values in the tree

  a
 / \
b   c
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type tree struct {
	value       int
	left, right *tree
}

func build(values []int) *tree {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	return root
}

func add(p *tree, v int) *tree {
	if p == nil {
		p = new(tree)
		p.value = v

		return p
	}

	if v < p.value {
		p.left = add(p.left, v)
	} else {
		p.right = add(p.right, v)
	}

	return p
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func (t *tree) String() string {
	var s string
	var values = make([]int, 0)
	values = appendValues(values, t)

	for i := range values {
		if i == 0 {
			s += fmt.Sprintf("[%v", values[i])
			continue
		}

		s += fmt.Sprintf(" %v", values[i])
	}

	s += "]"

	return s
}

func main() {
	min, max := 5, 10
	rand.Seed(time.Now().Unix())

	n := rand.Intn(max - min + 1) + min

	data := make([]int, n)
	for i := 0; i < n; i++ {
		data[i] = rand.Intn(100)
	}

	t := build(data)
	fmt.Println(t)
}
