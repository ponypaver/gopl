package main

// xkcd

//The popular web comic xkcd has a JSON interface. For example, a request to
//https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
//many favorites. Dow nlo ad each URL (once!) and bui ld an offline index. Write a tool xkcd
//that, using this index, prints the URL and transcript of each comic that matches a search term
//provided on the command line.

//{
// "month": "1",
// "num": 1,
// "link": "",
// "year": "2006",
// "news": "",
// "safe_title": "Barrel - Part 1",
// "transcript": "[[A boy sits in a barrel which is floating in an ocean.]]\nBoy: I wonder where I'll float next?\n[[The barrel drifts into the distance. Nothing else can be seen.]]\n{{Alt: Don't we all.}}",
// "alt": "Don't we all.",
// "img": "https://imgs.xkcd.com/comics/barrel_cropped_(1).jpg",
// "title": "Barrel - Part 1",
// "day": "1"
//}

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type cache struct {
	comics              []*comic
	total, currentTotal int
	startID             int
	missingNum          []int
	needUpdate          []int
	currentComicUrl     string
	urlTmpl             string
	cacheFile           string
}

type comic struct {
	Num        int    `json:"num"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	//SafeTitle  string `json:"safe_title"`
	//Alt        string `json:"alt"`
	//Day        string `json:"day"`
	//Img        string `json:"img"`
	//Link       string `json:"link"`
	//Month      string `json:"month"`
	//News       string `json:"news"`
	//Year       string `json:"year"`
}

func (c *cache) index() error {
	id := c.startID
	//for id := 0; id < c.currentTotal - len(c.missingNum); id++ {
	for {
		url := fmt.Sprintf(c.urlTmpl, id)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Errorf("get url: %v failed with err: %v\n", url, err)
		}

		if resp.StatusCode == http.StatusNotFound && in(id, c.missingNum) {
			fmt.Fprintf(os.Stderr, "comic with id: %v missing from xkcd.com for no reason. skipping\n", id)
			id++
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("get url: %v returned:%v, not 200 OK", url, resp.StatusCode)
		}

		//c.comics[id] = new(comic)
		if err := json.NewDecoder(resp.Body).Decode(&c.comics[id]); err != nil {
		//b, _ := ioutil.ReadAll(resp.Body)
		//if err := json.Unmarshal(b, &c.comics[id]); err != nil {
			resp.Body.Close()
			return fmt.Errorf("decode comic url: %v failed with err: %v\n", url, err)
		}
		resp.Body.Close()
		fmt.Printf("index comic number: %v done.\n", id)

		if id == c.currentTotal {
			break
		}

		id++
	}

	return nil
}

func (c *cache) setCurrentTotal() error {
	url := c.currentComicUrl
	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("get url: %v failed with err: %v\n", url, err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get url: %v returned:%v, not 200 OK", url, resp.StatusCode)
	}

	var re comic
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		resp.Body.Close()
		return fmt.Errorf("decode comic url: %v failed with err: %v\n", url, err)
	}
	resp.Body.Close()

	//c.total = re.Num
	c.currentTotal = re.Num
	return nil
}

func (c *cache) update() {

}

func (c *cache) save() error {
	file, err := os.Create(c.cacheFile)
	if err != nil {
		return err
	}
	defer file.Close()

	b, err := json.Marshal(c.comics)
	if _, err := file.Write(b); err != nil {
		return err
	}

	return nil
}

func (c *cache) load() error {
	if !fileExist(c.cacheFile) {
		fmt.Printf("No local cache found, will cache to: %v\n", c.cacheFile)

		if err := c.setCurrentTotal(); err != nil {
			return err
		}

		c.comics = make([]*comic, c.currentTotal+1)
		if err := c.index(); err != nil {
			return err
		}

		c.save()

		return nil
	}

	fmt.Printf("load cache from: %v\n", c.cacheFile)
	b, err := ioutil.ReadFile(c.cacheFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &c.comics); err != nil {
		return err
	}

	c.total = c.comics[len(c.comics)-1].Num

	return nil
}

func (c *cache) search(query string) []*comic {
	var re []*comic

	for i := 0; i < len(c.comics); i++ {
		if in(i, c.missingNum) {
			continue
		}
		if strings.Contains(c.comics[i].Title, query) || strings.Contains(c.comics[i].Transcript, query) {
			re = append(re, c.comics[i])
		}
	}

	return re
}

func fileExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func in(id int, ids []int) bool {
	for i := range ids {
		if id == ids[i] {
			return true
		}
	}

	return false
}

func newCache() *cache {
	return &cache{
		startID:         1,
		missingNum:      []int{0, 404},
		currentComicUrl: "https://xkcd.com/info.0.json",
		urlTmpl:         "https://xkcd.com/%v/info.0.json",
		cacheFile:       "xkcdComics.json",
	}
}

var query *string

func init() {
	query = flag.String("search", "demo", "keyword to search xkcd comic cache")
	flag.Parse()
}

func main() {
	cache := newCache()
	if err := cache.load(); err != nil {
		fmt.Fprintf(os.Stderr, "load cache failed: %v\n", err)
		os.Exit(1)
	}

	results := cache.search(*query)

	if len(results) > 0 {
		fmt.Println("\nsearch result:\nTitle\t\tTranscript")
		for _, re := range results {
			fmt.Println(re.Title, re.Transcript)
		}

		return
	}

	fmt.Fprintf(os.Stderr, "nothing found in cache for %q\n", *query)
	os.Exit(1)
}
