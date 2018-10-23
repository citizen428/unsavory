package main

import (
	"github.com/citizen428/unsavory/internal/config"
	"github.com/citizen428/unsavory/internal/unsavory"
)

func main() {
	u := unsavory.NewClient(config.Token, config.DryRun)
	u.Run()
}
