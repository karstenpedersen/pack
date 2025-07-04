package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetPathBaseAndExtension(path string) (string, string, string) {
	dirPath := filepath.Dir(path)
	base := filepath.Base(path)
	extension := strings.TrimPrefix(filepath.Ext(path), ".")
	return dirPath, base, extension
}

func FindFileInParents(filename string) (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		filePath := filepath.Join(current, filename)
		if _, err := os.Stat(filePath); err == nil {
			return filePath, nil
		}

		parent := filepath.Dir(current)
		if parent == current {
			break
		}
		current = parent
	}

	return "", fmt.Errorf("file %s not found in any parent directory", filename)
}
