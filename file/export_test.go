package file

import (
	"errors"
	"fmt"
	"os"
)

type Exported struct{}

var Export Exported

func (e Exported) TargetDir(name string) bool {
	return targetDir(name)
}

type fsStatErr struct {
	fsImpl
}

func (f fsStatErr) Stat(name string) (os.FileInfo, error) {
	return nil, errors.New(fmt.Sprint("Stat error: ", name))
}
func (e Exported) FSStatErr() func() {
	return swapFS(fsStatErr{})
}

type fsReadDirErr struct {
	fsImpl
}

func (f fsReadDirErr) ReadDir(dirname string) ([]os.FileInfo, error) {
	return nil, errors.New(fmt.Sprint("ReadDir error: ", dirname))
}
func (e Exported) FSReadDirErr() func() {
	return swapFS(fsReadDirErr{})
}

func swapFS(dummy fileSystem) func() {
	tmp := fs
	fs = dummy
	return func() {
		fs = tmp
	}
}
