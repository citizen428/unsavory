package unsavory

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func RunCLI() {
	var dryRun bool
	var proxyURL, token string

	flag.BoolVar(&dryRun, "dry-run", false, "Enables dry run mode")
	flag.StringVar(&proxyURL, "proxy-url", "", "HTTP proxy URL")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] token\n\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
	token = flag.Arg(0)

	if proxyURL != "" {
		os.Setenv("HTTP_PROXY", proxyURL)
	}

	log.SetFlags(0)

	httpClient := http.Client{
		Timeout: time.Second * 15,
	}
	NewClient(token, dryRun, &httpClient).Run()
}
