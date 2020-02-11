package util

import "runtime"

// IsWindows は実行環境がWindowsであるかどうかを返す
func IsWindows() bool {
	return runtime.GOOS == "windows"
}
