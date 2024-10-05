package pack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/karstenpedersen/pack/utils"
)

type ProjectHooks struct {
	PreHook  string `mapstructure:"preHook" json:"preHook,omitempty"`
	PostHook string `mapstructure:"postHook" json:"postHook,omitempty"`
}

type Project struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	OutDir  string            `json:"outDir"`
	Include []string          `json:"include"`
	Exclude []string          `json:"exclude,omitempty"`
	Rename  map[string]string `json:"rename,omitempty"`
	Hooks   ProjectHooks      `json:"hooks,omitempty"`
}

func DefaultProject(appConfig *App) (*Project, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	name := filepath.Base(wd)

	return &Project{
		Name:    name,
		Method:  appConfig.Method,
		OutDir:  appConfig.OutDir,
		Include: make([]string, 0),
		Exclude: make([]string, 0),
		Rename:  make(map[string]string),
		Hooks: ProjectHooks{
			PreHook:  "",
			PostHook: "",
		},
	}, nil
}

func LoadProject(app *App, path string, name string, typee string) (*Project, error) {
	configPath, err := GetProjectConfigPath()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	project, err := DefaultProject(app)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, project)

	return project, nil
}

func (p *Project) GetAffectedFiles() []string {
	return utils.GlobMatch(".", p.Include, p.Exclude)
}

func (p *Project) GetOutputPath() string {
	return p.OutDir
}

func (p *Project) GetTargetPath() string {
	outFile := fmt.Sprintf("%s.%s", p.Name, methodExtensions[p.Method])
	return path.Join(p.OutDir, outFile)
}

func (p *Project) Marshal() ([]byte, error) {
	return json.MarshalIndent(p, "", "    ")
}

func (p *Project) Pack() (string, error) {
	// Get files to package
	files := p.GetAffectedFiles()

	path := p.GetTargetPath()

	// Package project
	if p.Method == "zip" {
		if err := utils.ZipFiles(files, path, p.Rename); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("invalid method")
	}

	return path, nil
}

func GetProjectConfigPath() (string, error) {
	path := DEFAULT_PROJECT_CONFIG_FILE
	return path, nil
}

func ProjectConfigExists() (bool, error) {
	configPath, err := GetProjectConfigPath()
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(configPath); err != nil {
		return false, err
	}

	return true, nil
}
