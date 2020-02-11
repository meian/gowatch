package path

import "strings"

// UnixPath はWindowsパスの場合にUnixパスに変換する
func UnixPath(path string) string {
	if strings.Contains(path, `\`) {
		return strings.ReplaceAll(path, `\`, "/")
	}
	return path
}
