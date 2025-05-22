package main

import (
	"log"
	"os"

	"github.com/pr-impact-analyzer/internal/analyzer"
	"github.com/pr-impact-analyzer/internal/cli"
)

func main() {
	app := cli.NewApp()
	if err := app.Execute(); err != nil {
		log.Printf("Error: %v\n", err)
		os.Exit(1)
	}
} 