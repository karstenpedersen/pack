package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/karstenpedersen/pack/utils"
)

func init() {
	initCmd.Flags().BoolP("yes", "y", false, "Yes to all")
	initCmd.Flags().Bool("override", false, "Override existing config")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use: "init",
	Short: "Initialize project",
	Annotations: map[string]string{
		"skipProjectConfig": "true",
	},
	Run: func(cmd *cobra.Command, args []string) {
		yesToAll, _ := cmd.Flags().GetBool("yes")
		override, _ := cmd.Flags().GetBool("override")

		configFile := PROJECT_CONFIG_FILE + ".json"

		// Check if config file already exists
		if _, err := os.Stat(configFile); !override && err == nil {
			fmt.Println("Config file already exists.")
			os.Exit(1)
		}

		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		name := filepath.Base(wd)
		method := "zip"

		if !yesToAll {
			name = utils.Input("Name", name)
			method = utils.Input("Method", method)
		}

		var defaultConfig = fmt.Sprintf(`{
	"name": "%s",
	"method": "%s"
}`, name, method)

		// Creating config
		err = os.WriteFile(configFile, []byte(defaultConfig), 0644)
		if err != nil {
			fmt.Println("Error creating config file:", err)
			os.Exit(1)
		}

		fmt.Println("Initialized project:")
		fmt.Println(defaultConfig)
	},
}