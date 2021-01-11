package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// The URL to access the BecauseOfProg API
const APIUrl = "https://api.becauseofprog.fr/v1"

// APIResult stores the result of an API call
type APIResult struct {
	// The data returned by the call (in this case one or many publications)
	Data []Publication `json:"data"`
}

// Publication stores metadata about a BecauseOfProg article
type Publication struct {
	// Partial URL tho the publication
	URL string `json:"url"`
	// Title of this article
	Title string `json:"title"`
	// Date and time of the publication
	Timestamp int `json:"timestamp"`
	// Full URL to the illustration
	Banner string `json:"banner"`
	// Short description of the publication content
	Description string `json:"description"`
	// Information about article's author
	Author User `json:"author"`
}

// User represents a registered user on the BecauseOfProg
type User struct {
	// Full name of the user
	Name string `json:"displayname"`
}

// SearchArticles performs a search into publications of the BecauseOfProg
func SearchArticles(search string) (result APIResult, err error) {
	resp, err := http.Get(fmt.Sprintf("%s/blog-posts?search=%s", APIUrl, search))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}
