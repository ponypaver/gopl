package main

import (
	"fmt"
	"net/http"
	"sort"
)

type StrSlice []string

func (s StrSlice) Len() int           { return len(s) }
func (s StrSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s StrSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []string{"abahello", "abc", "abd"}
	sort.Strings(s)
	fmt.Println(s)
}