/*
Develop startElement and endElement into a general HTML pretty-printer.
Print comment nodes, text nodes, and the attributes of each element (<a href='...'>). Use
short for ms like <img/> instead of <img></img> when an element has no children. Write a
test to ensure that the output can be parsed successfully. (See Chapter 11.)
*/

package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"strings"
)

var depth int

func startElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
		depth++
	case html.TextNode:
		if strings.TrimSpace(n.Data) != "" {
			fmt.Printf("%*s%s\n", depth*2+2, "", n.Data)
			depth++
		}
	case html.ElementNode:
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s=%q", attr.Key, attr.Val)
		}
		if n.FirstChild != nil {
			fmt.Println(">")
		} else {
			fmt.Println("/>")
		}
		depth++
	}
}

func endElement(n *html.Node) {
	switch n.Type {
	case html.CommentNode:
		//	fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		depth--
	case html.TextNode:
		if strings.TrimSpace(n.Data) != "" {
			depth--
		}
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func forEachNode(node *html.Node, pre, post func(node *html.Node)) {
	if node != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	post(node)
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	forEachNode(root, startElement, endElement)
}
