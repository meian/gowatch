package path_test

import (
	"testing"

	"github.com/meian/gowatch/path"
	"github.com/stretchr/testify/assert"
)

func TestIsGoFile(t *testing.T) {
	a := assert.New(t)
	tests := []struct {
		desc     string
		file     string
		expected bool
	}{
		{desc: "blank string", file: "", expected: false},
		{desc: "not go file", file: "sample.js", expected: false},
		{desc: "go file", file: "sample.go", expected: true},
		{desc: "go test file", file: "sample_test.go", expected: true},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			actual := path.IsGoFile(tt.file)
			a.Equal(tt.expected, actual)
		})
	}
}
