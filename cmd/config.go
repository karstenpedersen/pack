package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)


func initGlobalConfig() {
	globalConfig = viper.New()
	globalConfig.SetDefault("test", "Hello, World!")

	globalConfig.SetConfigName("pack")
	globalConfig.SetConfigType("json")
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	globalConfig.AddConfigPath(home)
	globalConfig.ReadInConfig();
}

func initProjectConfig() {
	projectConfig = viper.New()
	projectConfig.SetDefault("include", "")
	projectConfig.SetDefault("exclude", "")
	projectConfig.SetDefault("outDir", ".pack/")
	projectConfig.SetDefault("tempDir", "$outDir/temp/")

	projectConfig.SetConfigName(PROJECT_CONFIG_FILE)
	globalConfig.SetConfigType("json")
	projectConfig.AddConfigPath(".")

	if err := projectConfig.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config file found.\nSetup project by running 'pack init'")
		} else {
			fmt.Println("Can't project config:", err)
		}
		os.Exit(1)
	}
}
