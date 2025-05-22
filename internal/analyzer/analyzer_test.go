package analyzer

import (
	"testing"
)

func TestIdentifyModule(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "simple module",
			filename: "internal/analyzer/analyzer.go",
			want:     "internal",
		},
		{
			name:     "root file",
			filename: "main.go",
			want:     "",
		},
		{
			name:     "nested module",
			filename: "pkg/api/v1/handlers.go",
			want:     "pkg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdentifyModule(tt.filename); got != tt.want {
				t.Errorf("IdentifyModule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRiskyChange(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		status   string
		want     bool
	}{
		{
			name:     "risky file",
			filename: "go.mod",
			status:   "modified",
			want:     true,
		},
		{
			name:     "safe file",
			filename: "README.md",
			status:   "modified",
			want:     false,
		},
		{
			name:     "deleted file",
			filename: "test.go",
			status:   "removed",
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRiskyChange(tt.filename, tt.status); got != tt.want {
				t.Errorf("IsRiskyChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
