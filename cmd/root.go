package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/karstenpedersen/pack/pack"
	"github.com/karstenpedersen/pack/utils"
)

var app *pack.App
var project *pack.Project
var projectConfigFile string

var rootCmd = &cobra.Command{
	Use:   "pack",
	Short: "Packages files",
	Long:  `Packages files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		app = pack.LoadApp()

		if _, skip := cmd.Annotations["skipProjectConfig"]; skip {
			return
		}

		fmt.Println("TEST", projectConfigFile)
		p, err := pack.LoadProject(app, projectConfigFile)
		if err != nil {
			utils.Exit(err)
		}
		project = p
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Create output directory
		os.MkdirAll(project.Config.OutDir, os.ModePerm)

		// Execute beforeHook
		if err := tryRunHook(project.Config.Hooks.PreHook); err != nil {
			utils.Exit("Error running preHook:", err)
		}

		path, err := project.Pack()
		if err != nil {
			utils.Exit(err)
		}

		// Execute afterHook
		if err := tryRunHook(project.Config.Hooks.PostHook); err != nil {
			utils.Exit("Error running postHook:", err)
		}

		fmt.Println(path)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.Exit("Error starting cli")
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&projectConfigFile, "config", "c", "", "config file")
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

func tryRunHook(hook string) error {
	if hook == "" {
		return nil
	}
	result, err := runHook(hook)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}
