package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/karstenpedersen/pack/pack"
	"github.com/karstenpedersen/pack/ui"
	"github.com/karstenpedersen/pack/utils"
	"github.com/spf13/cobra"
)

func init() {
	initCmd.Flags().BoolP("yes", "y", false, "Yes to all")
	initCmd.Flags().Bool("override", false, "Override existing config")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize pack config",
	Annotations: map[string]string{
		"skipProjectConfig": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		yesToAll, _ := cmd.Flags().GetBool("yes")
		override, _ := cmd.Flags().GetBool("override")

		configPath := pack.PROJECT_CONFIG_FILE
		absConfigPath, err := filepath.Abs(configPath)
		if err != nil {
			utils.Exit(err)
		}

		// Check if config file already exists
		if _, err := os.Stat(configPath); err == nil && !override {
			utils.Exit("Config file already exists.")
		}

		// Create default project
		project, err := pack.NewProject(app)
		if err != nil {
			utils.Exit(err)
		}

		// Get input from user
		if !yesToAll {
			ui.Input("Name", &project.Config.Name)
			ui.Input("Method", &project.Config.Method)
			ui.Input("Output directory", &project.Config.OutDir)
		}

		// Marshal config
		configData, err := project.MarshalConfig()
		if err != nil {
			utils.Exit("Failed to marshal config")
		}

		// Ask if config is OK
		if !yesToAll {
			// Show config to user
			fmt.Printf("About to write the following to %s:\n\n", absConfigPath)
			fmt.Println(string(configData))

			isThisOkay := "yes"
			ui.Input("\nIs this OK?", &isThisOkay)
			if isThisOkay != "yes" {
				fmt.Println("Aborted.")
				return
			}
		}

		// Create config file
		err = os.WriteFile(configPath, configData, 0644)
		if err != nil {
			utils.Exit("Error creating config file:", err)
		} else if yesToAll {
			// Show config to user
			fmt.Printf("Wrote the following to %s:\n\n", absConfigPath)
			fmt.Println(string(configData))
		}
	},
}
