package pack

import "fmt"

const (
	APP_NAME = "pack"
	VERSION  = "0.0.1"

	APP_CONFIG_FILE     = "config.json"
	PROJECT_CONFIG_FILE = "pack.config.json"
)

const (
	ZIP_METHOD = "zip"
	TAR_METHOD = "tar"
)

var METHOD_EXTENSIONS = map[string]string{
	ZIP_METHOD: "zip",
	TAR_METHOD: "tar",
}

const (
	DEFAULT_METHOD     string = ZIP_METHOD
	DEFAULT_OUTPUT_DIR string = ".pack-out"
)

func GetVersionString() string {
	return fmt.Sprintf("%s v%s", APP_NAME, VERSION)
}
