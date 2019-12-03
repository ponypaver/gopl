package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

func count(n *html.Node, re map[string]int) {

	if n.Type == html.ElementNode {
		re[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count(c, re)
	}
}

func main() {
	re := make(map[string]int)
	root, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	count(root, re)

	fmt.Printf("tag\tcount\n")
	for k, v := range re {
		fmt.Printf("%v\t%v\n", k, v)
	}
}
