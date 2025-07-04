package utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

func ReadConfigFile(v any, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	extension := filepath.Ext(path)
	if extension == ".json" {
		if err := json.Unmarshal(data, v); err != nil {
			return err
		}
	} else if extension == ".yaml" || extension == ".yml" {
		if err := yaml.Unmarshal(data, v); err != nil {
			return err
		}
	} else if extension == ".toml" {
		if err := toml.Unmarshal(data, v); err != nil {
			return err
		}
	}

	return nil
}
