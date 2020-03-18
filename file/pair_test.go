package file_test

import (
	p "path"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNewPair(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	tests := []struct {
		desc     string
		detected string
		test     string
		isError  bool
		enable   bool
	}{
		{desc: "no go file", detected: "src.js", test: "", isError: true},
		{desc: "src with test file", detected: "src.go", test: "src_test.go", enable: true},
		{desc: "test file", detected: "src_test.go", test: "src_test.go", enable: true},
		{desc: "no test file", detected: "no_test_src.go", test: "no_test_src_test.go", enable: false},
		{desc: "test file only", detected: "only_test.go", test: "only_test.go", enable: true},
		{desc: "is directory", detected: "dir.go", test: "dir_test.go", isError: true},
	}
	dir := "../testdata/pairtest"
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			dt := p.Join(dir, tt.detected)
			ts := p.Join(dir, tt.test)
			actual, err := file.NewPair(dt)
			if tt.isError {
				a.Error(err)
				return
			}
			a.NoError(err)
			a.Equal(dt, actual.Detected)
			a.Equal(ts, actual.Test)
			a.Equal(tt.enable, actual.TestEnabled())
		})
	}
}

func TestPairString(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	src := "../testdata/pairtest/src.go"
	pair, err := file.NewPair(src)
	a.NoError(err)
	a.Greater(len(pair.Detected), 0, "detected")
	a.Greater(len(pair.Test), 0, "test")
	a.NotEqual(pair.Detected, pair.Test)
	s := pair.String()
	a.Contains(s, pair.Detected)
	a.Contains(s, pair.Test)
}
