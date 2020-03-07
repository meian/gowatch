package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/stretchr/testify/assert"
)

func TestErrorMsg(t *testing.T) {
	a := assert.New(t)
	p := "fakeogr/afajejao/faef"
	err := path.NoGoFileError{Path: p}
	a.Error(err)
	a.Contains(err.Error(), p)
}
