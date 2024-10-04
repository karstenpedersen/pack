package utils

import (
	"io/fs"
	"path/filepath"
)

func Glob(root string, fn func(string)bool) []string {
	var files []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if fn(s) {
			files = append(files, s)
		}
		return nil
	})
	return files
}

func GlobMatch(root string, include []string, exclude []string) []string {
	return Glob(root, func (s string) bool {
		for _, i := range include {
			if matched, _ := filepath.Match(i, s); matched {
				for _, e := range exclude {
					if matched, _ := filepath.Match(e, s); matched {
						return false
					}
				}
				return true
			}
		}
		return false
	})
}