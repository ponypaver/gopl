/*
Without changing its behavior, rewrite the fetch function to use defer to close
the writable file.
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	fmt.Println("returned", local, resp.Request.URL)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		n, err = io.Copy(f, resp.Body)
		// Close file, but prefer error from Copy, if any.
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	return local, n, err
}

func main() {
	url := os.Args[1]
	local, n, err := fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s => %s (%d bytes).\n", url, local, n)
}
