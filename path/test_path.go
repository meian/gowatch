package path

import "strings"

const (
	// SrcSuffix はgoソースファイルのサフィックス
	SrcSuffix = ".go"
	// TestSuffix はテストソースファイルのサフィックス
	TestSuffix = "_test.go"
)

// ToTestPath はテストソースのパス(xxx_test.go)を返す
func ToTestPath(path string) (string, error) {
	switch {
	case !strings.HasSuffix(path, SrcSuffix):
		return "", NoGoFileError{path}
	case !strings.HasSuffix(path, TestSuffix):
		path = strings.TrimSuffix(path, SrcSuffix) + TestSuffix
	}
	return UnixPath(path), nil
}
