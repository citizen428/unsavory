package main

import (
	"flag"
	"fmt"
	"os"

	pb "github.com/citizen428/unsavory/internal/pinboard"
)

var token = flag.String("token", "", "Pinboard API token")

func main() {
	flag.Parse()
	if *token == "" {
		fmt.Fprintf(os.Stderr, "missing required API token\n")
		os.Exit(1)
	}

	client := pb.NewClient(*token)
	urls := client.GetAllURLs()
	results := make(chan pb.CheckResponse)
	fmt.Printf("Retrieved %d URLS\n", len(urls))

	for _, url := range urls {
		go client.CheckURL(url, results)
	}

	var result pb.CheckResponse
	for range urls {
		result = <-results
		if result.StatusCode == 404 {
			fmt.Printf("Deleting %s\n", result.URL)
			client.DeleteURL(result.URL)
		}
	}
}
