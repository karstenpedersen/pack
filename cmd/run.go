package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/karstenpedersen/pack/utils"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Packages files",
	Long:  `Packages files.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create output directory
		if err := os.MkdirAll(project.Config.OutDir, os.ModePerm); err != nil {
			utils.Exit("error: creating directory")
		}

		// Execute beforeHook
		if err := tryRunHook(project.Config.Hooks.PreHook); err != nil {
			utils.Exit("error: running preHook")
		}

		// Pack files
		path, err := project.Pack()
		if err != nil {
			utils.Exit(err)
		}

		// Execute afterHook
		if err := tryRunHook(project.Config.Hooks.PostHook); err != nil {
			utils.Exit("error: running postHook")
		}

		fmt.Println(path)
	},
}
