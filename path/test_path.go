package path

import "strings"

const (
	// テストソースファイルのサフィックス
	testSuffix = "_test.go"
)

// ToTestPath はテストソースのパス(xxx_test.go)を返す
func ToTestPath(path string) string {
	switch {
	case !IsGoFile(path):
		return path
	case strings.HasSuffix(path, testSuffix):
		return path
	}
	return strings.TrimSuffix(path, goSuffix) + testSuffix
}
