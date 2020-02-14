package test

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

func (cmd *Cmd) view() {
	log.Println("run test:", cmd.viewMsg())
}

func (cmd *Cmd) viewMsg() string {
	return fmt.Sprint(cmd)
}
