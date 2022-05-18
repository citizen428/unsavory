// Package unsavory provides the types, functions and methods for interacting
// with the Pinboard API and checking URLs for 404 status codes.
package unsavory

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	baseURL     = "https://api.pinboard.in/v1"
	userAgent   = "UnsavoryNG"
	workerCount = 50
)

// Client bundles all values necessary for API requests.
type Client struct {
	client *http.Client
	dryRun bool
	token  string
}

// NewClient returns a configured unsavory.Client.
func NewClient(token string, dryRun bool) *Client {
	client := &http.Client{
		Timeout: time.Second * 15,
	}

	return &Client{
		token:  token,
		dryRun: dryRun,
		client: client}
}

// Run fetches all URLs and kicks off the check process.
func (c *Client) Run() {
	if c.dryRun {
		log.Printf("You are using dry run mode. No links will be deleted!\n\n")
	}

	log.Println("Retrieving URLs")
	urls := c.getURLs()

	log.Printf("Retrieved %d URLS\n", len(urls))
	c.checkURLs(urls)
}

func (c *Client) getURLs() []string {
	var posts []struct {
		URL string `json:"href"`
	}

	resp := c.request("/posts/all")
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Could not retrieve URLs!\nPlease check your API token.")
	}

	json.NewDecoder(resp.Body).Decode(&posts)

	count := len(posts)
	urls := make([]string, count)

	for i, post := range posts {
		urls[i] = post.URL
	}
	return urls
}

func (c *Client) checkURLs(urls []string) {
	var wg sync.WaitGroup
	wg.Add(workerCount)

	ch := make(chan string)
	for i := 0; i < workerCount; i++ {
		go func(urls chan string) {
			defer wg.Done()
			for {
				u, ok := <-urls
				if !ok {
					return
				}

				c.checkURL(u)
			}
		}(ch)
	}

	for _, url := range urls {
		ch <- url
	}
	close(ch)

	wg.Wait()
}

func (c *Client) checkURL(u string) {
	resp, err := c.client.Head(u)
	if err != nil {
		if _, ok := err.(*url.Error); ok {
			log.Printf("Error (%s): %s\n", u, err)
		}
	} else {
		switch resp.StatusCode {
		case http.StatusOK:
			// do nothing
		case http.StatusNotFound, http.StatusGone:
			log.Printf("Deleting (404): %s\n", u)
			c.deleteURL(u)
		default:
			log.Printf("%d: %s\n", resp.StatusCode, u)
		}
	}
}

func (c *Client) deleteURL(url string) {
	if !c.dryRun {
		c.request("/posts/delete", url)
	}
}

func (c *Client) request(path string, query ...string) *http.Response {
	base, _ := url.Parse(baseURL)
	base.Path += path

	// Query params
	params := url.Values{}
	params.Add("format", "json")
	params.Add("auth_token", c.token)
	if len(query) > 0 {
		params.Add("url", query[0])
	}
	base.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", base.String(), nil)
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
