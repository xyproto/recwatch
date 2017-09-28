// Package recwatch provides a way to watch directories recursively with fsnotify

// +build gccgo

package recwatch

import (
	"os"
	"path/filepath"
	"strings"
)

type RecursiveWatcher struct {
	_       interface{}
	Files   chan string
	Folders chan string
}

func NewRecursiveWatcher(path string) (*RecursiveWatcher, error) {
	return nil, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string) error {
	return nil
}

// Subfolders returns a slice of subfolders (recursive), including the folder provided.
func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			if ShouldIgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

// ShouldIgnoreFile determines if a file should be ignored.
// File names that begin with "." or "_" are ignored by the go tool.
func ShouldIgnoreFile(name string) bool {
	return strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")
}
