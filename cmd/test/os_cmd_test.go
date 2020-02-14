package test_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/meian/gowatch/cmd/test"
	"github.com/stretchr/testify/assert"
)

var e = test.Export

func TestNewOsCmd(t *testing.T) {
	a := assert.New(t)
	cmd := e.NewCommand("go", "test")
	a.Equal(os.Stdout, cmd.Stdout)
	a.Equal(os.Stderr, cmd.Stderr)
}

func TestViewMsg(t *testing.T) {
	a := assert.New(t)
	cmd := e.NewCommand("go", "test")
	msg := e.CmdMsg(cmd)
	path, err := exec.LookPath("go")
	if err != nil {
		t.SkipNow()
	}
	a.Equal(fmt.Sprint(path, " test"), msg)
}
