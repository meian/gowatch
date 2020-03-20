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
		{desc: "not found", dirPath: "../testdata/noexists", exists: false},
		{desc: "not directory", dirPath: "../testdata/dirtest/a/b/c/dummy", exists: true},
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
	dirPath := "../testdata/dirtest"
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

func TestTargetDirs(t *testing.T) {
	testutil.ChCurrentDir()
	type args struct {
		dirPath   string
		recursive bool
	}
	tests := []struct {
		desc    string
		args    args
		want    []string
		wantErr bool
	}{
		{desc: "not exists - no recursive",
			args:    args{dirPath: "../testdata/noexists", recursive: false},
			wantErr: true},
		{desc: "not exists - recursive",
			args:    args{dirPath: "../testdata/noexists", recursive: true},
			wantErr: true},
		{desc: "no directory - no recursive",
			args:    args{dirPath: "../testdata/dirtest/a/b/c/dummy", recursive: false},
			wantErr: true},
		{desc: "no directory - recursive",
			args:    args{dirPath: "../testdata/dirtest/a/b/c/dummy", recursive: true},
			wantErr: true},
		{desc: "directory - no recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: false},
			want: []string{"../testdata/dirtest"}},
		{desc: "directory - recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: true},
			want: []string{
				"../testdata/dirtest",
				"../testdata/dirtest/a",
				"../testdata/dirtest/a/b",
				"../testdata/dirtest/a/b/c",
				"../testdata/dirtest/a/d"}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			got, err := file.TargetDirs(tt.args.dirPath, tt.args.recursive)
			a.Equal(tt.wantErr, err != nil)
			a.ElementsMatch(tt.want, got, got)
		})
	}
}

func TestTargetDirsStatErr(t *testing.T) {
	defer e.FSStatErr()()
	testutil.ChCurrentDir()
	type args struct {
		dirPath   string
		recursive bool
	}
	tests := []struct {
		desc string
		args args
	}{
		{desc: "directory - no recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: false}},
		{desc: "directory - recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: true}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			_, err := file.TargetDirs(tt.args.dirPath, tt.args.recursive)
			a.Error(err)
		})
	}
}

func TestTargetDirsReadDirErr(t *testing.T) {
	defer e.FSReadDirErr()()
	testutil.ChCurrentDir()
	type args struct {
		dirPath   string
		recursive bool
	}
	tests := []struct {
		desc    string
		args    args
		wantErr bool
	}{
		{desc: "directory - no recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: false}},
		{desc: "directory - recursive",
			args: args{dirPath: "../testdata/dirtest", recursive: true}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			got, err := file.TargetDirs(tt.args.dirPath, tt.args.recursive)
			a.NoError(err)
			a.NotNil(got)
		})
	}
}

func TestTargetDirPrivate(t *testing.T) {
	testutil.ChCurrentDir()
	tests := []struct {
		desc     string
		name     string
		expected bool
	}{
		{desc: "no dot name", name: "foo", expected: true},
		{desc: "current dir", name: ".", expected: true},
		{desc: "parent dir", name: "..", expected: true},
		{desc: "start with dot name", name: ".foo", expected: false},
		{desc: "contains dot name", name: "foo.bar", expected: true},
		{desc: "end with dot name", name: "foo.", expected: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			actual := e.TargetDir(tt.name)
			a.Equal(tt.expected, actual)
		})
	}
}
