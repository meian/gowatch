package file_test

import (
	p "path"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/stretchr/testify/assert"
)

func TestNewPair(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		detected string
		test     string
		isError  bool
		enable   bool
	}{
		{detected: "src.js", test: "", isError: true},
		{detected: "src.go", test: "src_test.go", enable: true},
		{detected: "src_test.go", test: "src_test.go", enable: true},
		{detected: "no_test_src.go", test: "no_test_src_test.go", enable: false},
		{detected: "only_test.go", test: "only_test.go", enable: true},
		{detected: "dir.go", test: "dir_test.go", isError: true}, // directory
	}
	dir := "../internal/pairtest"
	for _, test := range tests {
		dt := p.Join(dir, test.detected)
		ts := p.Join(dir, test.test)
		actual, err := file.NewPair(dt)
		if test.isError {
			a.Error(err, test)
			continue
		}
		a.NoError(err, test)
		a.Equal(dt, actual.Detected, test)
		a.Equal(ts, actual.Test, test)
		a.Equal(test.enable, actual.TestEnabled(), test)
	}
}
