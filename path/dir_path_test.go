package path_test

import (
	"fmt"
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/meian/gowatch/util"
	"github.com/stretchr/testify/assert"
)

func TestDirPath(t *testing.T) {
	tests := []struct {
		desc      string
		src       string
		dirIsUnix string
		dirIsWin  string
	}{
		{
			desc:      "normal path",
			src:       "foo/bar.go",
			dirIsUnix: "./foo",
			dirIsWin:  "./foo",
		},
		{
			desc:      "win normal path",
			src:       `foo\bar.go`,
			dirIsUnix: `./foo`,
			dirIsWin:  `./foo`,
		},
		{
			desc:      "dot start path",
			src:       "./foo/bar.go",
			dirIsUnix: "./foo",
			dirIsWin:  "./foo",
		},
		{
			desc:      "double dot start path",
			src:       "../foo/bar.go",
			dirIsUnix: "../foo",
			dirIsWin:  "../foo",
		},
		{
			desc:      "unix full path",
			src:       "/foo/bar.go",
			dirIsUnix: "/foo",
			dirIsWin:  "/foo",
		},
		{
			desc:      "win full path",
			src:       `c:\foo\bar.go`,
			dirIsUnix: `./c:/foo`, // no occurred
			dirIsWin:  `c:/foo`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			util.IsWindows = false
			dir := path.DirPath(tt.src)
			a.Equal(tt.dirIsUnix, dir, fmt.Sprintln("unix", tt.src))
			util.IsWindows = true
			dir = path.DirPath(tt.src)
			a.Equal(tt.dirIsWin, dir, fmt.Sprintln("windows", tt.src))
		})
	}
}
