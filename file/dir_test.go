package file_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

var e = file.Export

func TestRecurseFailed(t *testing.T) {
	testutil.ChCurrentDir()
	tests := []struct {
		desc    string
		dirPath string
		exists  bool
	}{
		{desc: "not found", dirPath: "../internal/noexists", exists: false},
		{desc: "not directory", dirPath: "../internal/dirtest/a/b/c/dummy", exists: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			_, err := file.RecurseDir(tt.dirPath)
			a.Error(err)
			_, err = os.Stat(tt.dirPath)
			a.Equal(tt.exists, err == nil)
		})
	}
}

func TestRecurseSuccess(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	dirPath := "../internal/dirtest"
	dirs, err := file.RecurseDir(dirPath)
	a.NoError(err)
	a.GreaterOrEqual(len(dirs), 3)
	pattern := fmt.Sprintf("^%v\\b", regexp.QuoteMeta(dirPath))
	for i, d := range dirs {
		t.Run(fmt.Sprintf("%d=>%s", i, d), func(t *testing.T) {
			a.DirExists(d)
			a.Regexp(pattern, d)
			a.NotRegexp(`/\.[^/]+\b`, d)
		})
	}
}

func TestContainsStartWithDot(t *testing.T) {
	testutil.ChCurrentDir()
	tests := []struct {
		desc     string
		name     string
		expected bool
	}{
		{desc: "no dot path", name: "foo/bar", expected: false},
		{desc: "last path is dot started", name: "foo/.bar", expected: true},
		{desc: "first path is dot started", name: ".foo/bar", expected: true},
		{desc: "middle path is dot started", name: "foo/.bar/sub", expected: true},
		{desc: "start path is curret", name: "./foo/bar", expected: false},
		{desc: "start path is parent", name: "../foo/bar", expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			actual := e.ContainsStartWithDot(tt.name)
			a.Equal(tt.expected, actual)
		})
	}
}

func TestTargetDirs(t *testing.T) {
	testutil.ChCurrentDir()
	type args struct {
		dirPath   string
		recursive bool
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "not exists - no recursive",
			args:    args{dirPath: "../internal/noexists", recursive: false},
			wantErr: true},
		{name: "not exists - recursive",
			args:    args{dirPath: "../internal/noexists", recursive: true},
			wantErr: true},
		{name: "no directory - no recursive",
			args:    args{dirPath: "../internal/dirtest/a/b/c/dummy", recursive: false},
			wantErr: true},
		{name: "no directory - recursive",
			args:    args{dirPath: "../internal/dirtest/a/b/c/dummy", recursive: true},
			wantErr: true},
		{name: "directory - no recursive",
			args: args{dirPath: "../internal/dirtest", recursive: false},
			want: []string{"../internal/dirtest"}},
		{name: "directory - recursive",
			args: args{dirPath: "../internal/dirtest", recursive: true},
			want: []string{
				"../internal/dirtest",
				"../internal/dirtest/a",
				"../internal/dirtest/a/b",
				"../internal/dirtest/a/b/c",
				"../internal/dirtest/a/d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := assert.New(t)
			got, err := file.TargetDirs(tt.args.dirPath, tt.args.recursive)
			a.Equal(tt.wantErr, err != nil)
			a.ElementsMatch(tt.want, got)
		})
	}
}
