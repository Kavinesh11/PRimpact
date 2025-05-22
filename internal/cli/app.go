package cli

import (
	"github.com/spf13/cobra"
)

type App struct {
	rootCmd *cobra.Command
}

func NewApp() *App {
	app := &App{
		rootCmd: &cobra.Command{
			Use:   "pr-impact-analyzer",
			Short: "Analyze the impact of pull requests",
			Long:  `A tool that analyzes the potential impact of pull requests by examining code changes, affected modules, and dependencies.`,
		},
	}

	app.addCommands()
	return app
}

func (a *App) Execute() error {
	return a.rootCmd.Execute()
}

func (a *App) addCommands() {
	analyzeCmd := &cobra.Command{
		Use:   "analyze",
		Short: "Analyze a pull request",
		RunE: func(cmd *cobra.Command, args []string) error {
			owner, _ := cmd.Flags().GetString("owner")
			repo, _ := cmd.Flags().GetString("repo")
			pr, _ := cmd.Flags().GetInt("pr")
			
			// TODO: Implement analysis logic
			return nil
		},
	}

	analyzeCmd.Flags().String("owner", "", "Repository owner")
	analyzeCmd.Flags().String("repo", "", "Repository name")
	analyzeCmd.Flags().Int("pr", 0, "Pull request number")
	analyzeCmd.MarkFlagRequired("owner")
	analyzeCmd.MarkFlagRequired("repo")
	analyzeCmd.MarkFlagRequired("pr")

	a.rootCmd.AddCommand(analyzeCmd)
} 