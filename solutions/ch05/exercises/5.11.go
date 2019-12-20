/*
The instructor of the linear algebra course decides that calculus is now a
prerequisite. Extend the topoSort function to report cycles.
*/

package main

import (
	"fmt"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func exist(item string, items []string) bool {
	for i := range items {
		if item == items[i] {
			return true
		}
	}

	return false
}

func transfer(start string, items []string, m map[string][]string) {
	if exist(start, items) {
		fmt.Printf("cycle detected, chain: %q\n", append(items, start))
		return
	}

	items = append(items, start)
	for _, item := range m[start] {
		transfer(item, items, m)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	var seen = make(map[string]bool)
	var visitAll func([]string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				transfer(item, nil, m)
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	for k := range m {
		visitAll([]string{k})
	}

	return order
}

func main() {
	var deps = map[string][]string{
		"a": {"b", "c"},
		"b": {"d"},
		"c": {"d"},
		"d": {"e", "f"},
		"e": {"y"},
		"f": {"b"},
		"x": {"i", "j"},
		"y": {"i", "z"},
	}
	//start := "a"
	//transfer(start, nil)
	for i, course := range topoSort(deps) {
		fmt.Printf("%d:\t%q\n", i+1, course)
	}

	fmt.Println()
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%q\n", i+1, course)
	}
}
