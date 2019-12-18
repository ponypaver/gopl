/*
Write a function
expand(s string, f func(string) string) string
that replaces each substring "$foo" within s by the text returned by f("foo").
*/

package main

import "strings"

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}
