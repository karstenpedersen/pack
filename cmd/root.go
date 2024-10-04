package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/karstenpedersen/pack/cfg"
	"github.com/karstenpedersen/pack/utils"
)


const APP_CONFIG_FILE = "config"
const PROJECT_CONFIG_FILE = ".pack"
var appConfig *viper.Viper = viper.New()
var projectConfig *viper.Viper = viper.New()
var projectConfigFile string

var rootCmd = &cobra.Command{
	Use:   "pack",
	Short: "Packages files",
	Long: `Packages files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := cfg.InitAppConfig(appConfig, APP_CONFIG_FILE)
		if err != nil {
			fmt.Println("Error reading app config")
			os.Exit(1)
		}

		if _, skip := cmd.Annotations["skipProjectConfig"]; skip {
			return
		}

		dir, name, ext := utils.GetPathBaseAndExtension(projectConfigFile)
		if name == "." {
			name = ".pack"
		}

		err = cfg.InitProjectConfig(projectConfig, dir, name, ext)
		if err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				fmt.Println("No config file found.\nSetup project by running 'pack init'")
			} else {
				fmt.Println("Can't load project config:", err)
			}
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		name := projectConfig.GetString("name")
		outDir := projectConfig.GetString("outDir")
		method := projectConfig.GetString("method")
		include := projectConfig.GetStringSlice("include")
		exclude := projectConfig.GetStringSlice("exclude")
		beforeHook := projectConfig.GetString("beforeHook")
		afterHook := projectConfig.GetString("afterHook")

		// Create output directory
		os.MkdirAll(outDir, os.ModePerm)

		// Execute beforeHook
		if beforeHook != "" {
			result, err := runHook(beforeHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("BEFORE HOOK:")
			fmt.Print(result)
		}

		// Get files to package
		files := utils.GlobMatch(".", include, exclude)

		outputPath := ""
		if (method == "zip") {
			outFile := name + ".zip"
			outputPath = filepath.Join(outDir, outFile)
			if err := utils.ZipFiles(files, outputPath); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			fmt.Println("Incorrect method:", method)
			os.Exit(1)
		}

		// Execute afterHook
		if afterHook != "" {
			result, err := runHook(afterHook)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println("AFTER HOOK:")
			fmt.Print(result)
		}

		fmt.Println(outputPath)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectConfigFile, "config", "c", ".", "config file")
}

func runHook(hook string) (string, error) {
	cmd := exec.Command(hook)
	var out strings.Builder
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}
