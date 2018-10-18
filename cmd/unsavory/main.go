package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	un "github.com/citizen428/unsavory/internal/unsavory"
)

var (
	dryRun   bool
	proxyURL string
	token    string
)

func init() {
	if cpu := runtime.NumCPU(); cpu == 1 {
		runtime.GOMAXPROCS(2)
	} else {
		runtime.GOMAXPROCS(cpu)
	}

	log.SetFlags(0)

	flag.BoolVar(&dryRun, "dry-run", false, "Enables dry run mode")
	flag.StringVar(&proxyURL, "proxy-url", "", "HTTP proxy URL")
	flag.StringVar(&token, "token", "", "Pinboard API token")
}

func main() {
	flag.Parse()
	if token == "" {
		log.Fatalln("Missing required API token")
	}

	if proxyURL != "" {
		os.Setenv("HTTP_PROXY", proxyURL)
	}

	un := un.NewClient(token, dryRun)
	un.Run()
}
