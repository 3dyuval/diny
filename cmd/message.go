package cmd

import (
	"fmt"
	"os"

	"github.com/dinoDanic/diny/commit"
	"github.com/dinoDanic/diny/config"
	"github.com/spf13/cobra"
)

// This file should only work in non-interactive environment
var messageCmd = &cobra.Command{
	Use:   "message",
	Short: "Generate commit message and output to stdout",
	Long: `Generate a commit message from staged changes and output to stdout.
Designed for piping to other commands or scripts.

Examples:
  diny message | git commit -F -
  diny message | pbcopy
  diny message > commit.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		gitDiff, err := commit.GetStagedDiff()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get git diff: %v\n", err)
			os.Exit(1)
		}

		if len(gitDiff) == 0 {
			fmt.Fprintf(os.Stderr, "No staged changes found. Stage files first with `git add`.\n")
			os.Exit(0)
		}

		diff := string(gitDiff)
		userConfig, err := config.Load()

		commitMessage, err := commit.CreateCommitMessage(diff, userConfig)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error generating commit message: %v\n", err)
			os.Exit(1)
		}

		fmt.Print(commitMessage)
	},
}

func init() {
	rootCmd.AddCommand(messageCmd)
}
