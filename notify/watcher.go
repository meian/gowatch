package notify

import "github.com/fsnotify/fsnotify"

// Watcher はfsnotify.Watcherの拡張型
type Watcher struct {
	*fsnotify.Watcher
	dirs map[string]bool
}

// NewWatcher はWatcherインスタンスを作成する
func NewWatcher() (*Watcher, error) {
	inWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	w := &Watcher{
		Watcher: inWatcher,
		dirs:    map[string]bool{},
	}
	return w, nil
}

// Add はディレクトリ監視を追加する、すでに追加されている場合は処理しない
func (w *Watcher) Add(dir string) error {
	if _, ok := w.dirs[dir]; ok {
		// already added
		return nil
	}
	err := w.Watcher.Add(dir)
	if err != nil {
		return err
	}
	w.dirs[dir] = true
	return nil
}

// Remove はディレクトリ監視を除外する、追加されていない場合は何もしない
func (w *Watcher) Remove(dir string) error {
	if _, ok := w.dirs[dir]; !ok {
		// not exists
		return nil
	}
	err := w.Watcher.Remove(dir)
	if err != nil {
		return err
	}
	delete(w.dirs, dir)
	return nil
}

// Watched はディレクトリが監視中のものであるかを返す
func (w *Watcher) Watched(dir string) bool {
	_, ok := w.dirs[dir]
	return ok
}
