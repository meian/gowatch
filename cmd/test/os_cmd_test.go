package test_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/meian/gowatch/cmd/test"
	"github.com/stretchr/testify/assert"
)

func TestNewCommand(t *testing.T) {
	a := assert.New(t)
	cmd := test.NewCommand("go", "test")
	a.Equal(os.Stdout, cmd.Stdout)
	a.Equal(os.Stderr, cmd.Stderr)
}

func TestCmdString(t *testing.T) {
	a := assert.New(t)
	cmd := test.NewCommand("go", "test")
	msg := cmd.String()
	path, err := exec.LookPath("go")
	if err != nil {
		t.Fatal("not found go executable")
	}
	a.Equal(path+" test", msg)
}
