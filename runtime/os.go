package runtime

import (
	r "runtime"
)

var (
	// IsWindows は実行環境がWindowsであるかどうかを返す、テスト時は値を後で設定可能
	IsWindows = false
)

func init() {
	IsWindows = r.GOOS == "windows"
}
