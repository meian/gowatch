package file

import "fmt"

// ReadError はファイルの読み込みに失敗した時のエラー
type ReadError struct {
	Name string
}

func (e ReadError) Error() string {
	return fmt.Sprint("cannot read file", e.Name)
}

// NoDirError はパスがディレクトリでない時のエラー
type NoDirError struct {
	Path string
}

func (e NoDirError) Error() string {
	return fmt.Sprintf("not directory: %v", e.Path)
}
