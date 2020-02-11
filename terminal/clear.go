package terminal

import (
	"os"
	"os/exec"

	"github.com/meian/gowatch/util"
)

// Clear はターミナルをクリアする
func Clear() error {
	cmd := clearCmd()
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// OS依存のターミナルクリアのコマンド
func clearCmd() *exec.Cmd {
	if util.IsWindows() {
		return exec.Command("cmd", "/c", "cls")
	}
	return exec.Command("clear")
}
