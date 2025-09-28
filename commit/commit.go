package commit

import (
	"fmt"
	"os"

	"github.com/dinoDanic/diny/config"
	"github.com/dinoDanic/diny/ui"
	"github.com/spf13/cobra"
)

func Main(cmd *cobra.Command, args []string) {
	fmt.Println()

	gitDiff, err := GetStagedDiff()

	if err != nil {
		ui.RenderError(fmt.Sprintf("Failed to get git diff: %v", err))
		os.Exit(1)
	}

	if len(gitDiff) == 0 {
		ui.RenderWarning("No staged changes found. Stage files first with `git add`.")
		os.Exit(0)
	}

	diff := string(gitDiff)

	userConfig, err := config.Load()

	var commitMessage string
	err = ui.WithSpinner("Generating your commit message...", func() error {
		var genErr error
		commitMessage, genErr = CreateCommitMessage(diff, userConfig)
		return genErr
	})

	if err != nil {
		ui.RenderError(fmt.Sprintf("%v", err))
		os.Exit(1)
	}

	HandleCommitFlow(commitMessage, diff, userConfig)
}
