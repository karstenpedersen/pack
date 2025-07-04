package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"

	"github.com/karstenpedersen/pack/pack"
	"github.com/karstenpedersen/pack/utils"
)

var app *pack.App
var project *pack.Project
var configPathFlag string

var rootCmd = &cobra.Command{
	Use:   "pack",
	Short: "Packages files",
	Long:  `Packages files.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		a, err := pack.LoadApp()
		if err != nil {
			utils.Exit(err)
		}
		app = a

		if skip, ok := cmd.Annotations["skipProjectConfig"]; ok && skip == "true" {
			return
		}

		p, err := pack.LoadCurrentProject(app, configPathFlag)
		if err != nil {
			utils.Exit(err)
		}
		project = p
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.Exit("error starting cli")
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPathFlag, "config", "c", "", "config file")
}

func runHook(hook string) (string, error) {
	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	cmd := exec.Command(hook)
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf
	if err := cmd.Run(); err != nil {
		return "", err
	}

	stdout := stdoutBuf.String()
	stderr := stderrBuf.String()
	if len(stderr) != 0 {
		return "", errors.New(stderr)
	}

	return stdout, nil
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
