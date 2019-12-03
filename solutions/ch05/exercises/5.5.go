/*
Implement countWordsAndImages. (See Exercise 4.9 for word-splitting.)
 */

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var (
	w, i int
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()

	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	switch {
	case n.Type == html.TextNode:
		words += countWords(n.Data)
	case n.Type == html.ElementNode && n.Data == "img":
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i = countWordsAndImages(c)
		words += w
		images += i
	}

	return
}

func countWords(s string) (words int) {
	text := bufio.NewScanner(strings.NewReader(s))
	text.Split(bufio.ScanWords)

	for text.Scan() {
		words++
	}

	return
}

func main() {
	fmt.Println(CountWordsAndImages(os.Args[1]))
}