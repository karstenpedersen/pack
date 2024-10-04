package cmd

import (
	"fmt"
	"os"
	"path"
	"archive/zip"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show files in pack.zip",
	Run: func(cmd *cobra.Command, args []string) {
		name := projectConfig.GetString("name")
		outDir := projectConfig.GetString("outDir")

		outFile := name + ".zip"
		outputPath := path.Join(outDir, outFile)

		zipListing, err := zip.OpenReader(outputPath)
		if err != nil {
			fmt.Println("run pack to package project")
			os.Exit(1)
		}
		defer zipListing.Close()
		for _, file := range zipListing.File {
			fmt.Println(file.Name)
		}
	},
}