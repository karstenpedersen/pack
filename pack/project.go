package pack

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/karstenpedersen/pack/utils"
)

type ProjectHooks struct {
	PreHook  string `json:"preHook,omitempty" yaml:"pre_hook,omitempty"`
	PostHook string `json:"postHook,omitempty" yaml:"post_hook,omitempty"`
}

type ProjectConfig struct {
	Name    string            `json:"name" yaml:"name"`
	Method  string            `json:"method" yaml:"method"`
	OutDir  string            `json:"outDir" yaml:"out_dir"`
	Include []string          `json:"include" yaml:"include"`
	Exclude []string          `json:"exclude,omitempty" yaml:"exclude,omitempty"`
	Rename  map[string]string `json:"rename,omitempty" yaml:"rename,omitempty"`
	Hooks   ProjectHooks      `json:"hooks,omitempty" yaml:"hooks,omitempty"`
}

type Project struct {
	Root       string
	ConfigPath string
	Config     ProjectConfig
}

func NewProject(app *App) (*Project, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Print(root)

	return DefaultProject(app, root)
}

func LoadProject(app *App, configPath string) (*Project, error) {
	project, err := DefaultProject(app, configPath)
	if err != nil {
		return nil, err
	}

	err = utils.ReadConfigFile(&project.Config, project.ConfigPath)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func DefaultProject(appConfig *App, configPath string) (*Project, error) {
	fileInfo, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}
	projectPath := filepath.Dir(configPath)
	if fileInfo.IsDir() {
		projectPath = configPath
	}
	projectName := filepath.Base(projectPath)

	return &Project{
		Root:       projectPath,
		ConfigPath: configPath,
		Config: ProjectConfig{
			Name:    projectName,
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

func LoadCurrentProject(app *App, optionalConfigPath ...string) (*Project, error) {
	configPath, err := GetProjectConfigPath(optionalConfigPath...)
	if err != nil {
		return nil, err
	}

	return LoadProject(app, configPath)
}

func (p *Project) GetConfig() ProjectConfig {
	return p.Config
}

func (p *Project) GetAffectedFiles() ([]string, error) {
	return utils.GlobMatch(".", p.Config.Include, p.Config.Exclude)
}

func (p *Project) GetMethodExtension() string {
	return METHOD_EXTENSIONS[p.Config.Method]
}

func (p *Project) GetTargetPath() string {
	extension := p.GetMethodExtension()
	filename := fmt.Sprintf("%s.%s", p.Config.Name, extension)
	return path.Join(p.Config.OutDir, filename)
}

func (p *Project) Marshal() ([]byte, error) {
	return json.MarshalIndent(p, "", "    ")
}

func (p *Project) MarshalConfig() ([]byte, error) {
	return json.MarshalIndent(p.Config, "", "    ")
}

func (p *Project) Pack() (string, error) {
	files, err := p.GetAffectedFiles()
	if err != nil {
		return "", err
	}
	path := p.GetTargetPath()

	// Package project
	switch p.Config.Method {
	case ZIP_METHOD:
		if err := utils.ZipFiles(files, path, p.Config.Rename); err != nil {
			return "", err
		}
	case TAR_METHOD:
		if err := utils.TarFiles(files, path, p.Config.Rename); err != nil {
			return "", err
		}
	default:
		return "", errors.New("invalid method")
	}

	return path, nil
}

func GetProjectConfigPath(optionalConfigPath ...string) (string, error) {
	if len(optionalConfigPath) == 1 && optionalConfigPath[0] != "" {
		return optionalConfigPath[0], nil
	}

	path, err := utils.FindFileInParents(PROJECT_CONFIG_FILE)
	if err != nil {
		return "", err
	}
	return path, nil
}
