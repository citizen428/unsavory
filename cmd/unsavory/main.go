package main

import (
	"github.com/citizen428/unsavory/internal/config"
	un "github.com/citizen428/unsavory/internal/unsavory"
)

func main() {
	un := un.NewClient(config.Token, config.DryRun)
	un.Run()
}
