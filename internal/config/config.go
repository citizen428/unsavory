package config

import (
	"flag"
	"log"
	"os"
	"runtime"
)

var (
	DryRun bool
	Token  string

	proxyURL string
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.SetFlags(0)

	flag.BoolVar(&DryRun, "dry-run", false, "Enables dry run mode")
	flag.StringVar(&Token, "token", "", "Pinboard API token")
	flag.StringVar(&proxyURL, "proxy-url", "", "HTTP proxy URL")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	if Token == "" {
		log.Fatalln("Missing required API token")
	}

	if proxyURL != "" {
		os.Setenv("HTTP_PROXY", proxyURL)
	}
}
