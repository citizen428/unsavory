package main

import (
	"flag"
	"log"

	un "github.com/citizen428/unsavory/internal/unsavory"
)

var (
	dryRun = flag.Bool("dry-run", false, "Enables dry run mode")
	token  = flag.String("token", "", "Pinboard API token")
)

func init() {
	log.SetFlags(0)
}

func main() {
	flag.Parse()
	if *token == "" {
		log.Fatalln("Missing required API token")
	}

	un := un.NewClient(*token, *dryRun)
	un.Run()
}
