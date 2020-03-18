package file_test

import (
	p "path"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	tests := []struct {
		desc     string
		path     string
		expected bool
	}{
		{desc: "is file", path: "a/b/c/dummy", expected: true},
		{desc: "is directory", path: "a/b/c", expected: false},
		{desc: "not exists", path: "a/b/c/dummy2", expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			fullPath := p.Join("../testdata/dirtest", tt.path)
			actual := file.IsFile(fullPath)
			a.Equal(tt.expected, actual)
		})
	}
}
