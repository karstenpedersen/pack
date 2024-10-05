package pack

import "fmt"

const (
	APP_NAME = "pack"
	VERSION  = "0.0.1"

	APP_CONFIG_FILE             = "config"
	DEFAULT_PROJECT_CONFIG_FILE = "pack.config.json"
)

const (
	zipMethod = "zip"
	tarMethod = "tar"
)

var methodExtensions = map[string]string{
	zipMethod: "zip",
	tarMethod: "tar",
}

const (
	defaultMethod string = zipMethod
	defaultOutDir string = ".pack-out"
)

func GetVersionString() string {
	return fmt.Sprintf("%s v%s", APP_NAME, VERSION)
}
