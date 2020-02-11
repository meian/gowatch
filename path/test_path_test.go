package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/stretchr/testify/assert"
)

func TestTestNameNoGoFile(t *testing.T) {
	a := assert.New(t)
	name := "sample.js"
	_, err := path.ToTestPath(name)
	a.Error(err)
	a.IsType(path.NoGoFileError{}, err)
	a.Contains(err.(path.NoGoFileError).Path, name)
}

func TestTestNameSrcPath(t *testing.T) {
	a := assert.New(t)
	name := "sample.go"
	actual, err := path.ToTestPath(name)
	a.NoError(err)
	a.Equal("sample_test.go", actual)
}

func TestTestNameTestPath(t *testing.T) {
	a := assert.New(t)
	name := "sample_test.go"
	actual, err := path.ToTestPath(name)
	a.NoError(err)
	a.Equal("sample_test.go", actual)
}

func TestTestNameMultiPath(t *testing.T) {
	a := assert.New(t)
	name := "a/b/c/sample.go"
	actual, err := path.ToTestPath(name)
	a.NoError(err)
	a.Equal("a/b/c/sample_test.go", actual)
}

func TestTestNameMultiPathWin(t *testing.T) {
	a := assert.New(t)
	name := `a\b\c\sample.go`
	actual, err := path.ToTestPath(name)
	a.NoError(err)
	a.Equal("a/b/c/sample_test.go", actual)
}

func TestTestNameComplexName(t *testing.T) {
	a := assert.New(t)
	name := "some_test.go.sample.go"
	actual, err := path.ToTestPath(name)
	a.NoError(err)
	a.Equal("some_test.go.sample_test.go", actual)
}
