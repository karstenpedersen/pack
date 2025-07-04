package pack

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/karstenpedersen/pack/utils"
)

type App struct {
	Method string `json:"method,omitempty"`
	OutDir string `json:"outDir,omitempty"`
}

func DefaultApp() *App {
	return &App{
		Method: DEFAULT_METHOD,
		OutDir: DEFAULT_OUTPUT_DIR,
	}
}

func LoadApp() (*App, error) {
	app := DefaultApp()

	configPath, err := GetAppConfigPath()
	if err != nil {
		return app, nil
	}

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		return app, nil
	}

	err = utils.ReadConfigFile(app, configPath)
	return nil, err
}

func GetAppConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, APP_NAME)

	return path, nil
}
