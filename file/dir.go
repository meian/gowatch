package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type fileSystem interface {
	Stat(name string) (os.FileInfo, error)
	ReadDir(dirname string) ([]os.FileInfo, error)
}
type fsImpl struct{}

var fs fileSystem = fsImpl{}

func (f fsImpl) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}
func (f fsImpl) ReadDir(dirname string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirname)
}

// TargetDirs は監視対象のディレクトリ一覧を返す、recursiveを指定する場合はサブディレクトリを含む
func TargetDirs(dirPath string, recursive bool) ([]string, error) {
	if dir, err := fs.Stat(dirPath); err != nil || !dir.IsDir() {
		return nil, NoDirError{Path: dirPath}
	}
	if recursive {
		return RecurseDir(dirPath)
	}
	return []string{dirPath}, nil
}

// RecurseDir は自身とサブディレクトリのパスリストを返す
func RecurseDir(dirPath string) ([]string, error) {
	st, err := fs.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !st.IsDir() {
		return nil, NoDirError{Path: dirPath}
	}
	bucket := &dirBucket{paths: []string{}}
	bucket.collect(dirPath)
	return bucket.paths, nil
}

// walk処理内でディレクトリのみを収集するコンテナ
type dirBucket struct {
	paths []string
}

// ディレクトリパスを再帰的に収集する
func (bucket *dirBucket) collect(dirPath string) {
	bucket.paths = append(bucket.paths, dirPath)
	fs, err := fs.ReadDir(dirPath)
	if err != nil {
		log.Println("[WARN]cannot read dir: ", dirPath)
		log.Println(err)
		return
	}
	for _, f := range fs {
		if !f.IsDir() {
			continue
		}
		name := f.Name()
		if !targetDir(name) {
			continue
		}
		bucket.collect(filepath.ToSlash(filepath.Join(dirPath, name)))
	}
}

func targetDir(name string) bool {
	switch {
	case name == ".":
		return true
	case name == "..":
		return true
	case strings.HasPrefix(name, "."):
		return false
	default:
		return true
	}
}
