package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


const GLOBAL_CONFIG_FILE = ".pack"
const PROJECT_CONFIG_FILE = ".pack"
var globalConfig *viper.Viper
var projectConfig *viper.Viper

var rootCmd = &cobra.Command{
	Use:   "pack",
	Short: "Packages files",
	Long: `Packages files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initGlobalConfig()

		if _, skip := cmd.Annotations["skipProjectConfig"]; skip {
			return
		}

		initProjectConfig()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	
}
