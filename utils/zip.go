package utils

import (
	"io"
	"os"
	"archive/zip"
)


func ZipFiles(files []string, path string) error {
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