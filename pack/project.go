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

type ProjectConfig struct {
	Name    string            `json:"name"`
	Method  string            `json:"method"`
	OutDir  string            `json:"outDir"`
	Include []string          `json:"include"`
	Exclude []string          `json:"exclude,omitempty"`
	Rename  map[string]string `json:"rename,omitempty"`
	Hooks   ProjectHooks      `json:"hooks,omitempty"`
}

type Project struct {
	Root   string
	Config ProjectConfig
}

func DefaultProject(appConfig *App) (*Project, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	name := filepath.Base(wd)

	return &Project{
		Root: wd,
		Config: ProjectConfig{
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
		},
	}, nil
}

func LoadProject(app *App, path ...string) (*Project, error) {
	configPath, err := GetProjectConfigPath(path...)
	if err != nil {
		return nil, err
	}
	fmt.Println(configPath)

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
	if err = json.Unmarshal(data, &project.Config); err != nil {
		return nil, err
	}

	return project, nil
}

func (p *Project) GetAffectedFiles() []string {
	return utils.GlobMatch(".", p.Config.Include, p.Config.Exclude)
}

func (p *Project) GetOutputPath() string {
	return p.Config.OutDir
}

func (p *Project) GetTargetPath() string {
	outFile := fmt.Sprintf("%s.%s", p.Config.Name, methodExtensions[p.Config.Method])
	return path.Join(p.Config.OutDir, outFile)
}

func (p *Project) Marshal() ([]byte, error) {
	return json.MarshalIndent(p, "", "    ")
}

func (p *Project) Pack() (string, error) {
	// Get files to package
	files := p.GetAffectedFiles()

	path := p.GetTargetPath()

	// Package project
	if p.Config.Method == "zip" {
		if err := utils.ZipFiles(files, path, p.Config.Rename); err != nil {
			return "", err
		}
	} else {
		return "", errors.New("invalid method")
	}

	return path, nil
}

func GetProjectConfigPath(optionalPath ...string) (string, error) {
	if len(optionalPath) > 0 && optionalPath[0] != "" {
		return optionalPath[0], nil
	}

	path, err := FindFileInParents(DEFAULT_PROJECT_CONFIG_FILE)
	if err != nil {
		return "", err
	}
	return path, nil
}

func FindFileInParents(filename string) (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		filePath := filepath.Join(current, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", fmt.Errorf("file %s not found in any parent directory", filename)
}
