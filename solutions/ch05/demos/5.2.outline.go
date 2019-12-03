package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	//"net/http"
	"os"
)

func outline2(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline2(stack, c)
	}
}

func outline(stack []string, root *html.Node) {
	if root == nil {
		return
	}

	if root.Type == html.ElementNode {
		stack = append(stack, root.Data)
		fmt.Println(stack)
	}

	outline(stack, root.FirstChild)
	outline(stack, root.NextSibling)
}

func main() {
	//url := os.Args[1]
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//root, err := html.Parse(resp.Body)
	root, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	outline(nil, root)
}