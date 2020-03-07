package file_test

import (
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/stretchr/testify/assert"
)

func TestNoDirErrorMsg(t *testing.T) {
	a := assert.New(t)
	p := "fakeogr/afajejao/faef"
	err := file.NoDirError{Path: p}
	a.Error(err)
	a.Contains(err.Error(), p)
}

func TestReadrrorMsg(t *testing.T) {
	a := assert.New(t)
	p := "fakeogr/afajejao/faef"
	err := file.ReadError{Path: p}
	a.Error(err)
	a.Contains(err.Error(), p)
}
