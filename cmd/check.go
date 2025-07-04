package cmd

import (
	"fmt"

	"github.com/karstenpedersen/pack/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check what files will be packaged",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := project.GetAffectedFiles()
		if err != nil {
			utils.Exit(err)
		}

		if len(files) == 0 {
			fmt.Println("No files matched include rules.")
			return
		}

		for _, file := range files {
			new, exists := project.Config.Rename[file]
			if exists {
				fmt.Printf("%s > %s\n", file, new)
			} else {
				fmt.Println(file)
			}
		}
	},
}
