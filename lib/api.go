package lib

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
	// The total number of pages (each page contains up to 10 entries)
	Pages int
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

func (p *Publication) FormatLink() (url string, link string) {
	url = fmt.Sprintf("https://becauseofprog.fr/article/%s", p.URL)
	link = fmt.Sprintf("[%s](%s)", p.Title, url)
	return
}

// User represents a registered user on the BecauseOfProg
type User struct {
	// Full name of the user
	Name string `json:"displayname"`
}

// GetPublicationsByCategory searches for publications that are contained in a specific category
func GetPublicationsByCategory(category string, page int) (result APIResult, err error) {
	body, err := MakeRequest(fmt.Sprintf("blog-posts?category=%s&page=%d", category, page))
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}

// GetPublicationsBySearch performs a search into publications of the BecauseOfProg
func GetPublicationsBySearch(search string) (result APIResult, err error) {
	body, err := MakeRequest("blog-posts?search=" + search)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &result)
	return
}

func MakeRequest(url string) (body []byte, err error) {
	var resp *http.Response
	resp, err = http.Get(APIUrl + "/" + url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}
