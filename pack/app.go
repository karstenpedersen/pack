package pack

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
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

	err = json.Unmarshal(data, app)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func GetAppConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, APP_NAME)

	return path, nil
}
