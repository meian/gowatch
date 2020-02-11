package path

import "fmt"

// NoGoFileError はgoソースのファイルでない時のエラー
type NoGoFileError struct {
	Path string
}

func (e NoGoFileError) Error() string {
	return fmt.Sprintf("not go file: %v", e.Path)
}
