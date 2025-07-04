package utils

import (
	"io/fs"
	"path/filepath"
)

func Glob(root string, fn func(string) (bool, error)) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if match, err := fn(s); err != nil {
			return err
		} else if match {
			files = append(files, s)
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

func GlobMatch(root string, include []string, exclude []string) ([]string, error) {
	return Glob(root, func(s string) (bool, error) {
		for _, includePattern := range include {
			if matched, err := filepath.Match(includePattern, s); err != nil {
				return false, err
			} else if matched {
				for _, excludePattern := range exclude {
					if matched, err := filepath.Match(excludePattern, s); err != nil {
						return false, err
					} else if matched {
						return matched, nil
					}
				}
				return true, nil
			}
		}
		return false, nil
	})
}
