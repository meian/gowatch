package notify

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
)

type Exported struct{}

var Export Exported

func newWatcherError() (*fsnotify.Watcher, error) {
	return nil, errors.New("new watcher error")
}
func (e Exported) MockNewWatcherError() func() {
	tmp := newNativeWatcher
	newNativeWatcher = newWatcherError
	return func() {
		newNativeWatcher = tmp
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
