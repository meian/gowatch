//+build windows

package terminal

import (
	"os/exec"
)

func clearCmd() *exec.Cmd {
	return exec.Command("cmd", "/c", "cls")
}
