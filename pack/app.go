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
		Method: defaultMethod,
		OutDir: defaultOutDir,
	}
}

func LoadApp() *App {
	app := DefaultApp()

	configPath, err := GetAppConfigPath()
	if err != nil {
		return app
	}

	file, err := os.Open(configPath)
	if err != nil {
		return app
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return app
	}

	json.Unmarshal(data, app)

	return app
}

func GetAppConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, APP_NAME)

	return path, nil
}
