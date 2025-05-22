package main

import (
	"log"
	"net/http"

	"github.com/pr-impact-analyzer/internal/config"
	"github.com/pr-impact-analyzer/internal/github"
)

func main() {
	// Load GitHub App configuration
	appConfig := config.DefaultGitHubAppConfig()
	app := github.NewApp(appConfig)

	// Set up routes
	http.HandleFunc("/webhook", app.HandleWebhook)
	http.HandleFunc("/install", app.InstallApp)
	http.HandleFunc("/callback", app.Callback)

	// Start server
	port := ":8080"
	log.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 