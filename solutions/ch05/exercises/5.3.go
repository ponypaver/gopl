package main

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"log"
	"net/http"
	"os"
	"unicode"
)

func printText(n *html.Node) {
	if n.Data == "script" || n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		if ! unicode.IsSpace(rune(n.Data[0])) {
			fmt.Printf("<%s> %s </%s>\n", n.Parent.Data, n.Data, n.Parent.Data)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	root, err := html.Parse(utf8Reader)
	if err != nil {
		log.Fatal(err)
	}

	printText(root)
}
