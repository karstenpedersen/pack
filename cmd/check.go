package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/karstenpedersen/pack/utils"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use: "check",
	Short: "Check what files will be packaged",
	Run: func(cmd *cobra.Command, args []string) {
		include := projectConfig.GetStringSlice("include")
		exclude := projectConfig.GetStringSlice("exclude")
		files := utils.GlobMatch(".", include, exclude)
		for _, file := range files {
			fmt.Println(file)
		}
	},
}