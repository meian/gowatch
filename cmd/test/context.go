package test

import (
	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/notify"
)

const (
	// None はテストが開始されていない状態
	None = 1 + iota
	// Waiting はテスト開始を待機している状態
	Waiting
	// Executing はテスト実行中の状態
	Executing
)

// Context はテスト実行時の情報を保持する
type Context struct {
	Config      *Config
	Watcher     *notify.Watcher
	Directories []string
	Changed     *file.PairMap
	State       int
	Triggered   bool
	Done        chan error
}

// NewContext はコンフィグから実行情報を生成する
func NewContext(config *Config) (*Context, error) {
	c := &Context{
		Config:  config,
		Changed: file.NewPairMap(),
		State:   None,
		Done:    make(chan error),
	}
	if config.Recursive {
		dirs, err := file.RecurseDir(config.Dir)
		if err != nil {
			return nil, err
		}
		c.Directories = dirs
	} else {
		c.Directories = []string{config.Dir}
	}
	return c, nil
}
