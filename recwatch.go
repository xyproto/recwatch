// Package recwatch provides a way to watch directories recursively with fsnotify

package recwatch

// +build !gccgo

import (
	"errors"

	"github.com/fsnotify/fsnotify"
)

type (
	Event fsnotify.Event

	RecursiveWatcher struct {
		*fsnotify.Watcher
		Files   chan string
		Folders chan string
	}
)

func (e Event) String() string {
	return fsnotify.Event(e).String()
}

func NewRecursiveWatcher(path string) (*RecursiveWatcher, error) {
	folders := Subfolders(path)
	if len(folders) == 0 {
		return nil, errors.New("No folders to watch.")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	rw := &RecursiveWatcher{Watcher: watcher}

	rw.Files = make(chan string, 10)
	rw.Folders = make(chan string, len(folders))

	for _, folder := range folders {
		if err = rw.AddFolder(folder); err != nil {
			return nil, err
		}
	}
	return rw, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string) error {
	if err := watcher.Add(folder); err != nil {
		return err
	}
	watcher.Folders <- folder
	return nil
}
