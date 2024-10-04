package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Print the version number",
	Annotations: map[string]string{
		"skipProjectConfig": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pack v0.1")
	},
}