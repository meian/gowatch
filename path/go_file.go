package path

import "strings"

const (
	// goソースファイルのサフィックス
	goSuffix = ".go"
)

// IsGoFile はファイル名がgoのソースであるかを返す
func IsGoFile(file string) bool {
	return strings.HasSuffix(file, goSuffix)
}
