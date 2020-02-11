package file

import (
	"os"
	"path/filepath"

	"github.com/meian/gowatch/path"
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
	bucket := &dirBucket{paths: make([]string, 0)}
	err = filepath.Walk(dirPath, bucket.walk)
	if err != nil {
		// 再帰関数内はエラーを返さないけどそもそもディレクトリが読めなかった対策
		return nil, err
	}
	return bucket.paths, nil
}

// walk処理内でディレクトリのみを収集するコンテナ
type dirBucket struct {
	paths []string
}

// walk内で見つかったディレクトリパスを格納する
func (bucket *dirBucket) walk(name string, file os.FileInfo, err error) error {
	if file.IsDir() {
		bucket.paths = append(bucket.paths, path.UnixPath(name))
	}
	return nil
}
