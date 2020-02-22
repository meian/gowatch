package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/stretchr/testify/assert"
)

func TestTestNameNoGoFile(t *testing.T) {
	a := assert.New(t)
	name := "sample.js"
	file := path.ToTestPath(name)
	a.Equal("sample.js", file)
}

func TestTestNameSrcPath(t *testing.T) {
	a := assert.New(t)
	name := "sample.go"
	actual := path.ToTestPath(name)
	a.Equal("sample_test.go", actual)
}

func TestTestNameTestPath(t *testing.T) {
	a := assert.New(t)
	name := "sample_test.go"
	actual := path.ToTestPath(name)
	a.Equal("sample_test.go", actual)
}

func TestTestNameMultiPath(t *testing.T) {
	a := assert.New(t)
	name := "a/b/c/sample.go"
	actual := path.ToTestPath(name)
	a.Equal("a/b/c/sample_test.go", actual)
}

func TestTestNameMultiPathWin(t *testing.T) {
	a := assert.New(t)
	name := `a\b\c\sample.go`
	actual := path.ToTestPath(name)
	a.Equal(`a\b\c\sample_test.go`, actual)
}

func TestTestNameComplexName(t *testing.T) {
	a := assert.New(t)
	name := "some_test.go.sample.go"
	actual := path.ToTestPath(name)
	a.Equal("some_test.go.sample_test.go", actual)
}
