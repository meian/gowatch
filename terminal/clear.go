package terminal

import (
	"os"
	"os/exec"

	"github.com/meian/gowatch/runtime"
)

var (
	clrCmd *exec.Cmd
)

func init() {
	if runtime.IsWindows {
		clrCmd = exec.Command("cmd", "/c", "cls")
	} else {
		clrCmd = exec.Command("clear")
	}
	clrCmd.Stdout = os.Stdout
}

// Clear はターミナルをクリアする
func Clear() error {
	// ANSI escapeでもクリアできるけどWindowsの標準ターミナルで動かないのでコマンドで処理する
	return clrCmd.Run()
}
