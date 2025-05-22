package analyzer

import (
	"strings"
)

// IdentifyModule extracts the module name from a file path
func IdentifyModule(filename string) string {
	if filename == "main.go" {
		return ""
	}

	// Split the path and return the first directory
	parts := strings.Split(filename, "/")
	if len(parts) > 1 {
		return parts[0]
	}
	return ""
}

// IsRiskyChange determines if a file change is considered risky
func IsRiskyChange(filename string, status string) bool {
	// Deleted files are always considered risky
	if status == "removed" {
		return true
	}

	// List of files that are considered risky when modified
	riskyFiles := map[string]bool{
		"go.mod": true,
		"go.sum": true,
	}

	return riskyFiles[filename]
}
