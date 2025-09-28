/*
Copyright © 2025 NAME HERE dino.danic@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/dinoDanic/diny/config"
	"github.com/dinoDanic/diny/ui"
	"github.com/spf13/cobra"
)

// RunConfigurationSetup runs the interactive configuration setup and returns the config
func RunConfigurationSetup() config.UserConfig {
	// Start with default configuration values
	userConfig := config.UserConfig{
		UseEmoji:        false,
		UseConventional: false,
		Tone:            config.Casual,
		Length:          config.Short,
	}

	// Emoji confirmation
	err := huh.NewConfirm().
		Title("Use emoji prefixes in commit messages?").
		Description("Add emojis like ✨ feat: or 🐛 fix: to commit messages").
		Affirmative("Yes").
		Negative("No").
		Value(&userConfig.UseEmoji).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Conventional commits confirmation
	err = huh.NewConfirm().
		Title("Use Conventional Commits format?").
		Description("Format: type(scope): description").
		Affirmative("Yes").
		Negative("No").
		Value(&userConfig.UseConventional).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Tone selection
	err = huh.NewSelect[config.Tone]().
		Title("Choose your commit message tone").
		Options(
			huh.NewOption("Professional - formal and matter-of-fact", config.Professional),
			huh.NewOption("Casual - light but clear", config.Casual),
			huh.NewOption("Friendly - warm and approachable", config.Friendly),
		).
		Value(&userConfig.Tone).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Length selection
	err = huh.NewSelect[config.Length]().
		Title("Choose your commit message length").
		Options(
			huh.NewOption("Short - subject only (no body)", config.Short),
			huh.NewOption("Normal - subject + optional body (1-4 bullets)", config.Normal),
			huh.NewOption("Long - subject + detailed body (2-6 bullets)", config.Long),
		).
		Value(&userConfig.Length).
		Run()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	return userConfig
}

type InitOption string

const (
	ConfigureProject InitOption = "configure"
	InstallHooks     InitOption = "hooks"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Diny for your project",
	Long: `Initialize Diny for your project with interactive setup options.

You can choose to:
- Configure project settings (commit message preferences)
- Install git hooks (auto-populate commit messages)

Configuration includes:
- Emoji: Add emoji prefixes to commit messages
- Format: Conventional commits or free-form messages
- Tone: Professional, casual, or friendly
- Length: Short, normal, or detailed messages

Git hooks will automatically populate commit messages using diny when you run 'git commit'.`,
	Run: func(cmd *cobra.Command, args []string) {
		var selectedOption InitOption

		// Ask user what they want to initialize
		err := huh.NewSelect[InitOption]().
			Title("What would you like to initialize?").
			Options(
				huh.NewOption("Configure project settings", ConfigureProject),
				huh.NewOption("Install git hooks", InstallHooks),
			).
			Value(&selectedOption).
			Run()

		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		// Execute based on selection
		switch selectedOption {
		case ConfigureProject:
			runConfigurationSetup()
		case InstallHooks:
			runHookInstallation()
		}
	},
}

func runConfigurationSetup() {
	ui.RenderTitle("🔧 Configuration Setup")
	userConfig := RunConfigurationSetup()

	err := config.Save(userConfig)
	if err != nil {
		fmt.Printf("Error saving configuration: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	config.PrintConfiguration(userConfig)
	ui.RenderTitle("✅ Configuration saved!")
}

func runHookInstallation() {
	ui.RenderTitle("🪝 Git Hook Installation")

	if err := installGitHook(); err != nil {
		fmt.Fprintf(os.Stderr, "Error installing git hook: %v\n", err)
		os.Exit(1)
	}

	ui.RenderTitle("✅ Git hook installed!")
	fmt.Println("Now when you run 'git commit', diny will pre-populate your commit message.")
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
