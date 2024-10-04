package cfg

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)


func InitAppConfig(config *viper.Viper, name string) error {
	config.SetConfigName(name)
	path, err := GetAppConfigPath();
	if err != nil {
		return nil
	}

	if _, err := os.Stat(path); err != nil {
		return nil
	}

	config.AddConfigPath(path)
	if err := config.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func GetAppConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	path = filepath.Join(path, "pack")

	return path, nil
}
