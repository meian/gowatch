package file

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// TargetDirs は監視対象のディレクトリ一覧を返す、recursiveを指定する場合はサブディレクトリを含む
func TargetDirs(dirPath string, recursive bool) ([]string, error) {
	if dir, err := os.Stat(dirPath); err != nil || !dir.IsDir() {
		return nil, NoDirError{Path: dirPath}
	}
	if recursive {
		return RecurseDir(dirPath)
	}
	return []string{dirPath}, nil
}

// RecurseDir は自身とサブディレクトリのパスリストを返す
func RecurseDir(dirPath string) ([]string, error) {
	st, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !st.IsDir() {
		return nil, NoDirError{Path: dirPath}
	}
	bucket := &dirBucket{paths: []string{}}
	if err := bucket.collect(dirPath); err != nil {
		return nil, err
	}
	return bucket.paths, nil
}

// walk処理内でディレクトリのみを収集するコンテナ
type dirBucket struct {
	paths []string
}

// ディレクトリパスを再帰的に収集する
func (bucket *dirBucket) collect(dirPath string) error {
	bucket.paths = append(bucket.paths, dirPath)
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
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
	return nil
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
