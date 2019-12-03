package main

//{
//"url": "https://api.github.com/repos/thedevsaddam/gojsonq/issues/67",
//"repository_url": "https://api.github.com/repos/thedevsaddam/gojsonq",
//"labels_url": "https://api.github.com/repos/thedevsaddam/gojsonq/issues/67/labels{/name}",
//"comments_url": "https://api.github.com/repos/thedevsaddam/gojsonq/issues/67/comments",
//"events_url": "https://api.github.com/repos/thedevsaddam/gojsonq/issues/67/events",
//"html_url": "https://github.com/thedevsaddam/gojsonq/issues/67",
//"id": 508728118,
//"node_id": "MDU6SXNzdWU1MDg3MjgxMTg=",
//"number": 67,
//"title": "FromInterface custom decoder",
//"user": {
//"login": "Dragomir-Ivanov",
//"id": 92195,
//"node_id": "MDQ6VXNlcjkyMTk1",
//"avatar_url": "https://avatars0.githubusercontent.com/u/92195?v=4",
//"gravatar_id": "",
//"url": "https://api.github.com/users/Dragomir-Ivanov",
//"html_url": "https://github.com/Dragomir-Ivanov",
//"followers_url": "https://api.github.com/users/Dragomir-Ivanov/followers",
//"following_url": "https://api.github.com/users/Dragomir-Ivanov/following{/other_user}",
//"gists_url": "https://api.github.com/users/Dragomir-Ivanov/gists{/gist_id}",
//"starred_url": "https://api.github.com/users/Dragomir-Ivanov/starred{/owner}{/repo}",
//"subscriptions_url": "https://api.github.com/users/Dragomir-Ivanov/subscriptions",
//"organizations_url": "https://api.github.com/users/Dragomir-Ivanov/orgs",
//"repos_url": "https://api.github.com/users/Dragomir-Ivanov/repos",
//"events_url": "https://api.github.com/users/Dragomir-Ivanov/events{/privacy}",
//"received_events_url": "https://api.github.com/users/Dragomir-Ivanov/received_events",
//"type": "User",
//"site_admin": false
//},
//"labels": [
//
//],
//"state": "open",
//"locked": false,
//"assignee": null,
//"assignees": [
//
//],
//"milestone": null,
//"comments": 1,
//"created_at": "2019-10-17T21:23:34Z",
//"updated_at": "2019-10-24T19:12:25Z",
//"closed_at": null,
//"author_association": "NONE",
//"body": "Hi there,\r\nIs there any way to load JSON document from `map[string]interface{}`?\r\nSimplest use case is querying MongoDB, which returns result documents in that format(without using schema struct). I am willing to make a PR if little help is supplied.\r\nThanks",
//"score": 81.48864
//},

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
	Items      []*Issue
}
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	log.Printf("get %v\n", IssuesURL+"?q="+q)
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

	result, err := SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
