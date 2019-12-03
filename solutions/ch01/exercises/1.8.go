package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	scheme := "http://"
	urls := os.Args[1:]

	for _, url := range urls {
		if ! strings.HasPrefix(url, scheme) {
			url = scheme + url
		}

		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch url: %v failed with err: %v\n", url, err)
			continue
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "copy body of url: %v failed with err: %v\n", url, err)
			continue
		}
	}
}
