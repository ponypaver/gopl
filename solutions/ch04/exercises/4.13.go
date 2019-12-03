package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// poster

//The JSON-based web service of the Open Movie Database lets you search
//https://omdbapi.com/ for a movie by name and download its poster image . Write a tool
//poster that downloads the poster image for the movie named on the command line.

const (
	apiTmpl = "http://www.omdbapi.com/?t=%v"
)

var name string

type movie struct {
	PosterUrl string `json:"Poster"`
}
func init() {
	flag.StringVar(&name, "name", "titanic", "move name")
	flag.Parse()
}

func getPosterUrl() (string, error) {
	movieUrl := fmt.Sprintf(apiTmpl, name)
	resp, err := http.Get(movieUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get %v returned none OK code: %v\n", movieUrl, resp.StatusCode)
	}

	var m movie
	if err := json.NewDecoder(resp.Body).Decode(m); err != nil {
		return "", err
	}

	return m.PosterUrl, nil

}

func download(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(name+".jpg", b, 0644); err != nil {
		return err
	}

	return nil
}

func main() {
	url, err := getPosterUrl()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	url = "https://m.media-amazon.com/images/M/MV5BMDdmZGU3NDQtY2E5My00ZTliLWIzOTUtMTY4ZGI1YjdiNjk3XkEyXkFqcGdeQXVyNTA4NzY1MzY@._V1_SX300.jpg"
	if err := download(url); err != nil {
		fmt.Fprintf(os.Stderr, "download poster for movie: %v failed: %v\n", name, err)
		os.Exit(1)
	}

	fmt.Printf("poster of movie: %v saved to %v\n", name, name+".jpg")
}