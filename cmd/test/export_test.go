package test

import "os/exec"

type Exported struct{}

var Export Exported

func (e Exported) NewCommand(path string, args ...string) *exec.Cmd {
	return newCommand(path, args...).Cmd
}

func (e Exported) CmdMsg(cmd *exec.Cmd) string {
	return (&Cmd{Cmd: cmd}).viewMsg()
}
