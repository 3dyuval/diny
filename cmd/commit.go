package cmd

import (
	"fmt"
	"github.com/dinoDanic/diny/helpers"
	"github.com/dinoDanic/diny/ollama"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "🦖 Generate AI-powered commit messages from staged changes",
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

		if cleanDiffLen > 2000 {
			fmt.Println("⚠️ Large changeset detected — this may take longer to process ⏳")
		}

		// fmt.Println(".............................ORGINAL..........................")
		// fmt.Println(gitDiff)
		// fmt.Println(".............................ORGINAL END..........................")
		// fmt.Printf("\n")
		// fmt.Printf("\n")
		//
		// fmt.Println(".............................CLEAN..........................")
		// fmt.Println(cleanDiff)
		// fmt.Println(".............................CLEAN END..........................")
		//
		// fmt.Printf("\n")
		// fmt.Printf("\n")
		fmt.Printf("📏 Diff size → Raw: %d chars | Cleaned: %d chars\n", gitDiffLen, cleanDiffLen)

		if cleanDiffLen == 0 {
			fmt.Println("🌱 No meaningful content detected in the diff.")
			os.Exit(0)
		}

		systemPrompt := slimdiff.BuildSystemPrompt()
		fullPrompt := systemPrompt + cleanDiff
		commitMessage, err := ollama.Main(fullPrompt)
		if err != nil {
			fmt.Printf("💥 Error generating commit message: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\n")
		fmt.Printf("\n")
		fmt.Printf("\n🦖 Suggested commit message:\n\n%s\n\n", commitMessage)

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
