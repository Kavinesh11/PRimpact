package config

// GitHubAppConfig holds the configuration for the GitHub App
type GitHubAppConfig struct {
	// App name and basic info
	Name        string
	Description string
	HomepageURL string

	// OAuth and Authorization
	CallbackURL string
	SetupURL    string

	// Webhook configuration
	WebhookURL string
	WebhookSecret string

	// Permissions
	RepositoryPermissions map[string]string
	OrganizationPermissions map[string]string
	AccountPermissions map[string]string

	// Events to subscribe to
	Events []string

	// Installation target
	InstallationTarget string
}

// DefaultGitHubAppConfig returns the default configuration for the PR Impact Analyzer GitHub App
func DefaultGitHubAppConfig() *GitHubAppConfig {
	baseURL := "https://abc123def456.ngrok.io" // Replace with your actual ngrok URL

	return &GitHubAppConfig{
		Name:        "PR Impact Analyzer",
		Description: "Analyzes the potential impact of pull requests by examining code changes, affected modules, and dependencies",
		HomepageURL: baseURL,

		// OAuth and Authorization
		CallbackURL: baseURL + "/callback",
		SetupURL:    baseURL + "/setup",

		// Webhook configuration
		WebhookURL:    baseURL + "/webhook",
		WebhookSecret: "", // To be set during app creation

		// Repository permissions
		RepositoryPermissions: map[string]string{
			"contents":      "read",
			"metadata":      "read",
			"pull_requests": "read",
			"issues":        "read",
		},

		// Organization permissions
		OrganizationPermissions: map[string]string{
			"members": "read",
		},

		// Account permissions
		AccountPermissions: map[string]string{
			"email": "read",
		},

		// Events to subscribe to
		Events: []string{
			"pull_request",
			"pull_request_review",
			"issues",
			"meta",
		},

		// Installation target
		InstallationTarget: "any_account",
	}
} 