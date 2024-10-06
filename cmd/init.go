package cmd

import (
	"fmt"
	"os"

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
	Short: "Initialize project",
	Annotations: map[string]string{
		"skipProjectConfig": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		yesToAll, _ := cmd.Flags().GetBool("yes")
		override, _ := cmd.Flags().GetBool("override")

		configPath, err := pack.GetProjectConfigPath()
		if err != nil {
			utils.Exit(err)
		}
		if _, err := os.Stat(configPath); err == nil && !override {
			utils.Exit("Config file already exists.")
		}

		// Default config
		config, err := pack.DefaultProject(app)
		if err != nil {
			utils.Exit(err)
		}

		// Get input from user
		if !yesToAll {
			ui.Input("Name", &config.Config.Name)
			ui.Input("Method", &config.Config.Method)
			ui.Input("Output directory", &config.Config.OutDir)
		}

		// Marshal config
		configStr, err := config.Marshal()
		if err != nil {
			utils.Exit("Failed to marshal config")
		}

		// Creating config
		err = os.WriteFile(configPath, []byte(configStr), 0644)
		if err != nil {
			utils.Exit("Error creating config file:", err)
		}

		fmt.Println("Initialized project:")
		fmt.Println(string(configStr))
	},
}
