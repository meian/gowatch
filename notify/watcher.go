package notify

import "github.com/fsnotify/fsnotify"

// 内部ウォッチャーを生成するインターフェイス
type watcherGen interface {
	newWatcher() (*fsnotify.Watcher, error)
}
type watcherGenImpl struct{}

func (gen watcherGenImpl) newWatcher() (*fsnotify.Watcher, error) {
	return fsnotify.NewWatcher()
}

var gen watcherGen = watcherGenImpl{}

// Watcher はfsnotify.Watcherの拡張型
type Watcher struct {
	*fsnotify.Watcher
	dirs   map[string]bool
	add    func(name string) error
	remove func(name string) error
}

// NewWatcher はWatcherインスタンスを作成する
func NewWatcher() (*Watcher, error) {
	watcher, err := gen.newWatcher()
	if err != nil {
		return nil, err
	}
	w := &Watcher{
		Watcher: watcher,
		dirs:    map[string]bool{},
		add:     watcher.Add,
		remove:  watcher.Remove,
	}
	return w, nil
}

// Add はディレクトリ監視を追加する
func (w *Watcher) Add(name string) error {
	if _, ok := w.dirs[name]; ok {
		// already added
		return nil
	}
	err := w.add(name)
	if err != nil {
		return err
	}
	w.dirs[name] = true
	return nil
}

// Remove はディレクトリ監視を除外する
func (w *Watcher) Remove(name string) error {
	if _, ok := w.dirs[name]; !ok {
		// not exists
		return nil
	}
	err := w.remove(name)
	if err != nil {
		return err
	}
	delete(w.dirs, name)
	return nil
}

// Close は監視処理を閉じる
func (w *Watcher) Close() error {
	err := w.Watcher.Close()
	w.dirs = map[string]bool{}
	return err
}

// Watched はディレクトリが監視中のものであるかを返す
func (w *Watcher) Watched(name string) bool {
	_, ok := w.dirs[name]
	return ok
}
