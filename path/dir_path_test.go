package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/meian/gowatch/runtime"
	"github.com/stretchr/testify/assert"
)

func TestDirPath(t *testing.T) {
	// WindowsパスのテストはWindows上のみで行う
	// UnixパスのテストはWindowsも含めて全環境で行う
	tests := []struct {
		desc   string
		src    string
		dir    string
		forWin bool
	}{
		{desc: "normal path",
			src: "foo/bar.go",
			dir: "./foo",
		},
		{desc: "win normal path",
			src:    `foo\bar.go`,
			dir:    `./foo`,
			forWin: true,
		},
		{desc: "dot start path",
			src: "./foo/bar.go",
			dir: "./foo",
		},
		{desc: "double dot start path",
			src: "../foo/bar.go",
			dir: "../foo",
		},
		{desc: "unix full path",
			src: "/foo/bar.go",
			dir: "/foo",
		},
		{desc: "win full path",
			src:    `c:\foo\bar.go`,
			dir:    `c:/foo`,
			forWin: true,
		},
		{desc: "no path",
			src: "bar.go",
			dir: ".",
		},
		{desc: "dot start no path",
			src: "./bar.go",
			dir: ".",
		},
		{desc: "double dot start no path",
			src: "../bar.go",
			dir: "..",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if !runtime.IsWindows && tt.forWin {
				t.SkipNow()
			}
			a := assert.New(t)
			dir := path.DirPath(tt.src)
			a.Equal(tt.dir, dir, tt.src)
		})
	}
}
