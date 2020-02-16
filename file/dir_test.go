package file_test

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/stretchr/testify/assert"
)

var e = file.Export

func TestRecurseNotFound(t *testing.T) {
	chdirCurrent(t)
	a := assert.New(t)
	dirPath := "../internal/noexists"
	_, err := file.RecurseDir(dirPath)
	a.True(os.IsNotExist(err))
}

func TestRecurseNoDir(t *testing.T) {
	chdirCurrent(t)
	a := assert.New(t)
	dirPath := "../internal/dirtest/a/b/c/dummy"
	_, err := file.RecurseDir(dirPath)
	a.Error(err)
}

func TestRecurseSuccess(t *testing.T) {
	chdirCurrent(t)
	a := assert.New(t)
	dirPath := "../internal/dirtest"
	dirs, err := file.RecurseDir(dirPath)
	a.NoError(err)
	a.GreaterOrEqual(len(dirs), 3)
	pattern := fmt.Sprintf("^%v\\b", regexp.QuoteMeta(dirPath))
	for i, d := range dirs {
		t.Log(i, d)
		a.DirExists(d)
		a.Regexp(pattern, d)
		a.NotRegexp(`/\.[^/]+\b`, d)
	}
}

func TestContainsStartWithDot(t *testing.T) {
	chdirCurrent(t)
	a := assert.New(t)
	tests := []struct {
		name     string
		expected bool
	}{
		{name: "foo/bar", expected: false},
		{name: "foo/.bar", expected: true},
		{name: "./foo/bar", expected: false},
		{name: "../foo/bar", expected: false},
		{name: "foo/.bar/sub", expected: true},
		{name: ".foo/bar", expected: true},
	}
	for _, test := range tests {
		actual := e.ContainsStartWithDot(test.name)
		a.Equal(test.expected, actual, test.name)
	}
}

func chdirCurrent(t *testing.T) {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("cannot get current path")
	}
	err := os.Chdir(filepath.Dir(file))
	if err != nil {
		t.Fatal(err)
		panic("cannot change directory")
	}
}
