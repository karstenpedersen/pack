package utils

import (
	"path/filepath"
	"strings"
)

func GetPathBaseAndExtension(path string) (string, string, string) {
	dirPath := filepath.Dir(path)
	base := filepath.Base(path)
	extension := strings.TrimPrefix(filepath.Ext(path), ".")
	return dirPath, base, extension
}
