package notify

import "github.com/fsnotify/fsnotify"

type fsWatcher interface {
	// Add はディレクトリ監視を追加する
	Add(name string) error
	// Remove はディレクトリ監視を除外する
	Remove(name string) error
	// Close は監視処理を閉じる
	Close() error
	// Events はファイルイベントを受け取るチャンネルを返す
	Events() chan fsnotify.Event
	// Errors はエラーを受け取るチャンネルを返す
	Errors() chan error
}

// Watcher はfsnotify.Watcherの拡張型
type Watcher interface {
	fsWatcher
	// Watched はディレクトリが監視中のものであるかを返す
	Watched(name string) bool
}

// 内部ウォッチャーを生成するインターフェイス
type watcherGen interface {
	newWatcher() (fsWatcher, error)
}
type watcherGenImpl struct{}

func (gen watcherGenImpl) newWatcher() (fsWatcher, error) {
	w, err := fsnotify.NewWatcher()
	return &nativeWatcher{w}, err
}

var gen watcherGen = watcherGenImpl{}

type nativeWatcher struct {
	*fsnotify.Watcher
}

func (w *nativeWatcher) Events() chan fsnotify.Event {
	return w.Watcher.Events
}
func (w *nativeWatcher) Errors() chan error {
	return w.Watcher.Errors
}

// Watcherの実装
type watcher struct {
	watcher fsWatcher
	dirs    map[string]bool
}

// NewWatcher はWatcherインスタンスを作成する
func NewWatcher() (Watcher, error) {
	inWatcher, err := gen.newWatcher()
	if err != nil {
		return nil, err
	}
	w := &watcher{
		watcher: inWatcher,
		dirs:    map[string]bool{},
	}
	return w, nil
}

func (w *watcher) Add(name string) error {
	if _, ok := w.dirs[name]; ok {
		// already added
		return nil
	}
	err := w.watcher.Add(name)
	if err != nil {
		return err
	}
	w.dirs[name] = true
	return nil
}

func (w *watcher) Remove(name string) error {
	if _, ok := w.dirs[name]; !ok {
		// not exists
		return nil
	}
	err := w.watcher.Remove(name)
	if err != nil {
		return err
	}
	delete(w.dirs, name)
	return nil
}

func (w *watcher) Close() error {
	err := w.watcher.Close()
	w.dirs = map[string]bool{}
	return err
}

func (w *watcher) Events() chan fsnotify.Event {
	return w.watcher.Events()
}

func (w *watcher) Errors() chan error {
	return w.watcher.Errors()
}

func (w *watcher) Watched(name string) bool {
	_, ok := w.dirs[name]
	return ok
}
