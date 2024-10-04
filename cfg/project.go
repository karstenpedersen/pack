package cfg

import (
	"github.com/spf13/viper"
)


func InitProjectConfig(config *viper.Viper, dir string, name string, ext string) error {
	config.SetConfigName(name)
	config.SetConfigType(ext)
	config.AddConfigPath(dir)

	config.SetDefault("include", "")
	config.SetDefault("exclude", "")
	config.SetDefault("method", "zip")
	config.SetDefault("outDir", ".pack-out/")
	config.SetDefault("beforeConfig", "")
	config.SetDefault("afterConfig", "")

	if err := config.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
