package test

import (
	"errors"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/meian/gowatch/file"
)

// LoopFSEvent はファイル監視イベントを処理するループ
func LoopFSEvent(c *Context) {
	for {
		select {
		case event, ok := <-c.Watcher.Events():
			if !ok {
				c.Done <- errors.New("cannot get fs event")
				return
			}
			switch {
			case opMatch(event, fsnotify.Write):
				onWrite(c, event)
			case opMatch(event, fsnotify.Create):
				onCreate(c, event)
			case opMatch(event, fsnotify.Remove):
				onRemove(c, event)
			}
		case err, ok := <-c.Watcher.Errors():
			if !ok {
				err = errors.New("cannot get fs error")
			}
			c.Done <- err
			return
		}
	}
}

// ファイルの書き込みがあった場合はテスト候補に追加
func onWrite(c *Context, event fsnotify.Event) {
	pair, err := file.NewPair(event.Name)
	if err != nil {
		// not a go file
		return
	}
	c.Changed.Add(pair)
	c.Triggered = true
}

// ディレクトリが作成された場合は監視対象に追加
func onCreate(c *Context, event fsnotify.Event) {
	dName := event.Name
	if stat, err := os.Stat(dName); err != nil || !stat.IsDir() {
		// not a directory
		return
	}
	log.Println("add new watch directory:", dName)
	c.Watcher.Add(dName)
}

// ディレクトリが削除された場合は監視対象から削除
func onRemove(c *Context, event fsnotify.Event) {
	dir := event.Name
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		// still exists
		return
	}
	if !c.Watcher.Watched(dir) {
		// not watched, include files.
		return
	}
	log.Println("remove watch directory:", dir)
	c.Watcher.Remove(dir)
}

func opMatch(e fsnotify.Event, op fsnotify.Op) bool {
	return e.Op&op == op
}
