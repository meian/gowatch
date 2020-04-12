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
	Watcher     *notify.Watcher
	Args        []string
	Directories []string
	Changed     *file.PairMap
	State       int
	Triggered   bool
	Done        chan error
}

// NewContext はコンフィグから実行情報を生成する
func NewContext(config *Config) (*Context, error) {
	dirs, err := collectDirs(config)
	if err != nil {
		return nil, err
	}
	return &Context{
		Changed:     file.NewPairMap(),
		State:       None,
		Done:        make(chan error),
		Args:        config.Args,
		Directories: dirs,
	}, nil
}

func collectDirs(config *Config) ([]string, error) {
	dirs := []string{}
	dirMap := map[string]bool{}
	for _, dir := range config.Dirs {
		tdirs, err := file.TargetDirs(dir, config.Recursive)
		if err != nil {
			return nil, err
		}
		for _, tdir := range tdirs {
			if _, ok := dirMap[tdir]; !ok {
				dirs = append(dirs, tdir)
				dirMap[tdir] = true
			}
		}
	}
	return dirs, nil
}
