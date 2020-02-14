package test

import (
	"os"
	"os/exec"
	"strings"
)

// Cmd はexec.Cmdのラッパー
type Cmd struct {
	*exec.Cmd
}

func newCommand(path string, args ...string) *Cmd {
	cmd := exec.Command(path, args...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	oc := &Cmd{Cmd: cmd}
	return oc
}

func (cmd *Cmd) viewMsg() string {
	params := append([]string{cmd.Path}, cmd.Args[1:]...)
	return strings.Join(params, " ")
}
