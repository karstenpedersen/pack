package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize project",
	Annotations: map[string]string{
		"skipProjectConfig": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		configFile := PROJECT_CONFIG_FILE + ".json"

		// Check if config file already exists
		if _, err := os.Stat(configFile); err == nil {
			fmt.Println("Config file already exists.")
			os.Exit(1)
		}

		var defaultConfig = []byte(`{
	"name": "assignment1",
	"include": [],
	"exclude": [],
	"outDir": "."
}
`)

		err := os.WriteFile(configFile, defaultConfig, 0644)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			os.Exit(1)
		}
		fmt.Println("Initialized project")
	},
}