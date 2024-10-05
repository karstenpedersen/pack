package utils

import (
	"archive/zip"
	"io"
	"os"
)

func ZipFiles(files []string, path string, renameRules map[string]string) error {
	archive, err := os.Create(path)
	if err != nil {
		return err
	}
	defer archive.Close()

	writer := zip.NewWriter(archive)
	defer writer.Close()

	// Copy files
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy file to archive
		value, exists := renameRules[filename]
		if exists {
			filename = value
		}
		write, err := writer.Create(filename)
		if err != nil {
			return err
		}
		if _, err := io.Copy(write, file); err != nil {
			return err
		}
	}

	return nil
}
