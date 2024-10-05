package cmd

import (
	"archive/zip"
	"fmt"

	"github.com/karstenpedersen/pack/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show files in output",
	Run: func(cmd *cobra.Command, args []string) {
		path := project.GetTargetPath()

		zipListing, err := zip.OpenReader(path)
		if err != nil {
			utils.Exit("run pack to package project:", err)
		}
		defer zipListing.Close()
		for _, file := range zipListing.File {
			fmt.Println(file.Name)
		}
	},
}
