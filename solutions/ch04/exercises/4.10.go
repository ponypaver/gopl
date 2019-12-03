package main
//Modify issues to report the results in age categories, say less than a month old,
//less than a year old, and more than a year old.

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}
type Issue struct {
	Number int
	HTMLURL string `json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"created_at"`
	Body string // in Markdown format
}
type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	log.Printf("get %v\n", IssuesURL + "?q=" + q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func main() {
	var lessThanAWeek, lessThanAMonth, lessThanAYear, moreThanAYear []*Issue

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range result.Items {
		days := time.Since(item.CreatedAt).Hours() / 24
		switch {
		case days <= 7:
			lessThanAWeek = append(lessThanAWeek, item)
		case days > 7 && days <= 30:
			lessThanAMonth = append(lessThanAMonth, item)
		case days > 30 && days <= 365:
			lessThanAYear = append(lessThanAYear, item)
		case days > 365:
			moreThanAYear = append(moreThanAYear, item)
		}
	}
	fmt.Printf("> %d issues less than a week:\n", len(lessThanAWeek))
	for _, item := range lessThanAWeek {
		fmt.Printf("#%-5d %9.9s  %.2f days ago %.55s\n",
			item.Number, item.User.Login, time.Since(item.CreatedAt).Hours() / 24, item.Title)
	}
	fmt.Println()
	fmt.Printf("> %d issues less than a Month:\n", len(lessThanAMonth))
	for _, item := range lessThanAMonth {
		fmt.Printf("#%-5d %9.9s  %.2f days ago %.55s\n",
			item.Number, item.User.Login, time.Since(item.CreatedAt).Hours() / 24, item.Title)
	}
	fmt.Println()
	fmt.Printf("> %d issues less than a year:\n", len(lessThanAYear))
	for _, item := range lessThanAYear {
		fmt.Printf("#%-5d %9.9s  %.2f days ago %.55s\n",
			item.Number, item.User.Login, time.Since(item.CreatedAt).Hours() / 24, item.Title)
	}
	fmt.Println()
	fmt.Printf("> %d issues more than a year:\n", len(moreThanAYear))
	for _, item := range moreThanAYear {
		fmt.Printf("#%-5d %9.9s  %.2f days ago %.55s\n",
			item.Number, item.User.Login, time.Since(item.CreatedAt).Hours() / 24, item.Title)
	}
}