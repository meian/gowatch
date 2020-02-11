package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/stretchr/testify/assert"
)

func TestUnixPathWin(t *testing.T) {
	a := assert.New(t)
	fPath := `sample\sample_func.go`
	actual := path.UnixPath(fPath)
	a.Equal("sample/sample_func.go", actual)
}

func TestUnixPathUnix(t *testing.T) {
	a := assert.New(t)
	fPath := "sample/sample_func.go"
	actual := path.UnixPath(fPath)
	a.Equal("sample/sample_func.go", actual)
}

func TestUnixPathNonDir(t *testing.T) {
	a := assert.New(t)
	fPath := "sample_func.go"
	actual := path.UnixPath(fPath)
	a.Equal("sample_func.go", actual)
}
