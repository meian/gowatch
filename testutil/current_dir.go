package testutil

import (
	"os"
	"path/filepath"
	"runtime"
)

// ChCurrentDir は呼び出し元ファイルの階層へ移動する
func ChCurrentDir() {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get current path")
	}
	err := os.Chdir(filepath.Dir(file))
	if err != nil {
		panic("cannot change directory")
	}
}
