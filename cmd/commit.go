package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/dinoDanic/diny/config"
	"github.com/dinoDanic/diny/helpers"
	"github.com/dinoDanic/diny/ollama"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate commit messages from staged changes",
	Long: `Diny analyzes your staged git changes and generates clear,
well-formatted commit messages using AI. 

This helps you keep a clean, consistent commit history with less effort.

Examples:
  diny commit
  diny commit --lang hr
  diny commit --style conventional`,
	Run: func(cmd *cobra.Command, args []string) {
		// Optimized git diff (ignores noise like lock files, binaries, etc.)
		gitDiffCmd := exec.Command("git", "diff", "--cached",
			"-U0", "--no-color", "--ignore-all-space", "--ignore-blank-lines",
			":(exclude)*.lock", ":(exclude)*package-lock.json", ":(exclude)*yarn.lock",
			":(exclude)node_modules/", ":(exclude)dist/", ":(exclude)build/")

		gitDiff, err := gitDiffCmd.Output()

		if err != nil {
			fmt.Printf("❌ Failed to get git diff: %v\n", err)
			os.Exit(1)
		}

		if len(gitDiff) == 0 {
			fmt.Println("🦴 No staged changes found. Stage files first with `git add`.")
			os.Exit(0)
		}

		cleanDiff := slimdiff.CleanForAI(string(gitDiff))
		gitDiffLen := len(gitDiff)
		cleanDiffLen := len(cleanDiff)

		fmt.Printf("📏 Diff size → Raw: %d chars | Cleaned: %d chars\n", gitDiffLen, cleanDiffLen)

		if cleanDiffLen > 2000 {
			fmt.Println("⚠️ Large changeset detected — this may take longer to process ⏳")
		}

		if cleanDiffLen == 0 {
			fmt.Println("🌱 No meaningful content detected in the diff.")
			os.Exit(0)
		}

		fmt.Print("🐢 My tiny server is thinking hard, hold tight! \n")

		userConfig := config.Load()

		systemPrompt := slimdiff.BuildSystemPrompt(userConfig)
		fullPrompt := systemPrompt + cleanDiff

		commitMessage, err := ollama.MainStream(fullPrompt)

		if err != nil {
			fmt.Printf("💥 Error generating commit message: %v\n", err)
			os.Exit(1)
		}

		if err != nil {
			fmt.Printf("Error displaying message: %v\n", err)
		}

		confirmed := confirmPrompt("👉 Do you want to commit with this message?")

		if confirmed {
			commitCmd := exec.Command("git", "commit", "--no-verify", "-m", commitMessage)
			err := commitCmd.Run()
			if err != nil {
				fmt.Printf("❌ Commit failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("✅ Commit successfully added to history!")
		} else {
			fmt.Println("🚫 Commit cancelled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
