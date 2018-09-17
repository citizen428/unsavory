// Package unsavory provides the types, functions and methods for interacting
// with the Pinboard API and checking URLs for 404 status codes.
package unsavory

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	baseURL   = "https://api.pinboard.in/v1"
	userAgent = "UnsavoryNG"
)

// Client bundles all values necessary for API requests.
type Client struct {
	baseQuery string
	client    *http.Client

	DryRun bool
	Token  string
}

type checkResponse struct {
	URL        string
	Status     string
	StatusCode int
}

// NewClient returns a configured unsavory.Client.
func NewClient(token string, dryRun bool) *Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		baseQuery: fmt.Sprintf("format=json&auth_token=%s", token),
		Token:     token,
		DryRun:    dryRun,
		client:    client}
}

// Run fetches all URLs, checks them individually and removes saved links return
//  a 404 status code.
func (c *Client) Run() {
	if c.DryRun {
		log.Printf("You are using dry run mode. No links will be deleted!\n\n")
	}

	log.Println("Retrieving URLs")
	urls := c.getAllURLs()
	log.Printf("Retrieved %d URLS\n", len(urls))

	results := make(chan checkResponse)
	for _, url := range urls {
		go c.checkURL(url, results)
	}

	var result checkResponse
	for range urls {
		result = <-results
		switch result.StatusCode {
		case 200:
			continue
		case 404:
			log.Printf("Deleting %s\n", result.URL)
			c.deleteURL(result.URL)
		default:
			log.Printf("%s:%s\n", result.Status, result.URL)
		}
	}
}

func (c *Client) getAllURLs() []string {
	var posts []struct {
		URL string `json:"href"`
	}

	resp := c.request("/posts/all")
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalln("Could not retrieve URLs!\nPlease check your API token.")
	}

	json.NewDecoder(resp.Body).Decode(&posts)

	urls := make([]string, len(posts))

	for i, post := range posts {
		urls[i] = post.URL
	}
	return urls
}

func (c *Client) checkURL(url string, results chan<- checkResponse) {
	resp, err := c.client.Head(url)
	if err != nil {
		results <- checkResponse{url, err.Error(), 0}
		return
	}
	results <- checkResponse{url, resp.Status, resp.StatusCode}
}

func (c *Client) deleteURL(url string) {
	if !c.DryRun {
		c.request("/posts/delete", fmt.Sprintf("url=%s", url))
	}
}

func (c *Client) request(path string, query ...string) *http.Response {
	url := fmt.Sprintf("%s%s?%s", baseURL, path, c.baseQuery)
	if len(query) > 0 {
		url = url + "&" + strings.Join(query, "&")
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}
