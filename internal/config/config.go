package config

import (
	"fmt"
	"os"
)

type Config struct {
	GitHubToken string
}

func Load() (*Config, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("GITHUB_TOKEN environment variable is required")
	}

	return &Config{
		GitHubToken: token,
	}, nil
} 