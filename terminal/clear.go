package terminal

import (
	"os"
	"os/exec"
)

var (
	clrCmd *exec.Cmd
)

func init() {
	clrCmd = clearCmd()
	clrCmd.Stdout = os.Stdout
}

// Clear はターミナルをクリアする
func Clear() error {
	// ANSI escapeでもクリアできるけどWindowsの標準ターミナルで動かないのでコマンドで処理する
	return clrCmd.Run()
}
