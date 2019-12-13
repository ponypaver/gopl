/*
Exercise 5.8: Modify forEachNode so that the pre and post functions return a boole result
indicating whether to continue the traversal. Use it to write a function ElementByID with the
The Go Programming Language following signature that finds the first HTML element with the specified id attribute.

The function should stop the traversal as soon as a match is found.
func ElementByID(doc *html.Node, id string) *html.Node
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func empty(node html.Node) bool {
	return node.Parent == nil &&
		node.FirstChild == nil &&
		node.LastChild == nil &&
		node.PrevSibling == nil &&
		node.NextSibling == nil
}

func forEachNode(node *html.Node, target *html.Node, pre, post func(node *html.Node) bool) {
	if pre != nil {
		if pre(node) {
			*target = *node
			return
		}
	}

	for c := node.FirstChild; c != nil && empty(*target); c = c.NextSibling {
		forEachNode(c, target, pre, post)
	}

	if post != nil {
		post(node)
	}
}

func ElementByID(doc *html.Node, id string) *html.Node {
	target := new(html.Node)
	pre := func(node *html.Node) bool {
		if node.Type == html.ElementNode {
			for i := range node.Attr {
				if node.Attr[i].Key == id {
					return true
				}
			}
		}
		return false
	}

	forEachNode(doc, target, pre, nil)
	return target
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

	id := "content"
	fmt.Println(ElementByID(root, id))
}
