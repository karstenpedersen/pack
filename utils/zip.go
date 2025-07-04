package utils

import (
	"archive/tar"
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

		// Rename file
		value, exists := renameRules[filename]
		if exists {
			filename = value
		}

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

func TarFiles(files []string, path string, renameRules map[string]string) error {
	archive, err := os.Create(path)
	if err != nil {
		return err
	}
	defer archive.Close()

	writer := tar.NewWriter(archive)
	defer writer.Close()

	// Copy files
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		fileInfo, err := os.Stat(filename)
		if err != nil {
			return err
		}

		value, exists := renameRules[filename]
		if exists {
			filename = value
		}

		// Copy file to archive
		hdr := &tar.Header{
			Name: filename,
			Mode: 0600,
			Size: fileInfo.Size(),
		}
		if err := writer.WriteHeader(hdr); err != nil {
			return err
		}
		if _, err := io.Copy(writer, file); err != nil {
			return err
		}
	}

	return nil
}
