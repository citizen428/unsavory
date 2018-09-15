package pinboard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type pinboard struct {
	baseURL   string
	userAgent string
	token     string

	client *http.Client
}

type CheckResponse struct {
	URL        string
	Status     string
	StatusCode int
}

func NewClient(token string) pinboard {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return pinboard{
		baseURL:   "https://api.pinboard.in/v1/",
		userAgent: "UnsavoryNG",
		token:     token,
		client:    client}
}

func (pb pinboard) GetAllURLs() []string {
	var posts []struct {
		URL string `json:"href"`
	}

	resp, err := pb.request("posts/all", "")
	if err != nil {
		fmt.Println("Could not retrieve URLs!\nPlease check your API token.")
		os.Exit(1)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&posts)

	urls := make([]string, len(posts))

	for i, post := range posts {
		urls[i] = post.URL
	}
	return urls
}

func (pb pinboard) CheckURL(url string, results chan<- CheckResponse) {
	resp, err := pb.client.Head(url)
	if err != nil {
		results <- CheckResponse{url, err.Error(), 0}
		return
	}
	results <- CheckResponse{url, resp.Status, resp.StatusCode}
}

func (pb pinboard) DeleteURL(url string) *http.Response {
	resp, err := pb.request("posts/delete", fmt.Sprintf("&url=%s", url))
	if err != nil {
		panic("Could not delete post")
	}
	return resp
}

func (pb pinboard) request(path string, query string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s?auth_token=%s&format=json%s", pb.baseURL, path, pb.token, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic("Invalid request")
	}
	return pb.client.Do(req)
}
