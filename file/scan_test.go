package file_test

import (
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	path := "../internal/srctest/test_for_scan.go"
	names, err := file.ScanTests(path)
	if !a.NoError(err) {
		a.FailNow("cannot scan")
	}
	t.Log(names)
	expected := []string{
		"TestSimpleTarget1",
		"TestSimpleTarget2",
		"TestSimpleTarget3",
		"TestHasAnounimous",
	}
	a.ElementsMatch(expected, names)
}
