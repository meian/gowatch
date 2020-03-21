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

func (gen watcherGenNewError) newWatcher() (*fsnotify.Watcher, error) {
	return nil, errors.New("new watcher error")
}
func (e Exported) WatcherGenNewError() func() {
	tmp := gen
	gen = watcherGenNewError{}
	return func() {
		gen = tmp
	}
}

func (e Exported) MockAddError(w *Watcher) {
	w.add = func(name string) error {
		return errors.New(fmt.Sprint("add error: ", name))
	}
}

func (e Exported) MockRemoveError(w *Watcher) {
	w.remove = func(name string) error {
		return errors.New(fmt.Sprint("remove error: ", name))
	}
}
