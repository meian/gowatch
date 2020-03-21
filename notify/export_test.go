package notify

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

type Exported struct{}

var Export Exported

type watcherGenNewError struct {
	watcherGenImpl
}

func (gen watcherGenNewError) newWatcher() (fsWatcher, error) {
	return nil, errors.New("new watcher error")
}
func (e Exported) WatcherGenNewError() func() {
	return swapGen(watcherGenNewError{})
}

type watcherGenAddError struct {
	watcherGenImpl
}
type nativeWatcherAddError struct {
	nativeWatcher
}

func (gen watcherGenAddError) newWatcher() (fsWatcher, error) {
	watcher, _ := fsnotify.NewWatcher()
	return &nativeWatcherAddError{nativeWatcher{watcher}}, nil
}
func (w *nativeWatcherAddError) Add(name string) error {
	return errors.New(fmt.Sprint("add error: ", name))
}
func (e Exported) WatcherGenAddError() func() {
	return swapGen(watcherGenAddError{})
}

type watcherGenRemoveError struct {
	watcherGenImpl
}
type nativeWatcherRemoveError struct {
	nativeWatcher
}

func (gen watcherGenRemoveError) newWatcher() (fsWatcher, error) {
	watcher, _ := fsnotify.NewWatcher()
	return &nativeWatcherRemoveError{nativeWatcher{watcher}}, nil
}
func (w *nativeWatcherRemoveError) Remove(name string) error {
	return errors.New(fmt.Sprint("remove error: ", name))
}
func (e Exported) WatcherGenRemoveError() func() {
	return swapGen(watcherGenRemoveError{})
}

func swapGen(g watcherGen) func() {
	tmp := gen
	gen = g
	return func() {
		gen = tmp
	}
}
