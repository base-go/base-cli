package main

import (
	"base-go/base-graph/commands"
	"os"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Exit(1)
	}
}
