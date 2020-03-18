package file_test

import (
	p "path"
	"testing"

	"github.com/meian/gowatch/file"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

func TestPairMap(t *testing.T) {
	testutil.ChCurrentDir()
	a := assert.New(t)
	tests := []struct {
		desc     string
		detected string
		has      bool
		count    int
	}{
		{desc: "src", detected: "src.go", has: false, count: 1},
		{desc: "src twice", detected: "src.go", has: true, count: 1},
		{desc: "test", detected: "src_test.go", has: false, count: 2},
		{desc: "only test", detected: "only_test.go", has: false, count: 3},
		{desc: "no test", detected: "no_test_src.go", has: false, count: 4},
	}
	m := file.NewPairMap()
	dir := "../testdata/pairtest"
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			d := p.Join(dir, tt.detected)
			p, err := file.NewPair(d)
			a.NoError(err)
			a.Equal(tt.has, m.Has(p.Detected))
			m.Add(p)
			a.Equal(m.Count(), tt.count)
		})
	}
	pMap, noTests := m.PopAll()
	a.Equal(m.Count(), 0)
	a.Len(noTests, 1)
	a.Len(pMap, 2)
	tc := 0
	for _, pss := range pMap {
		tc += len(pss)
	}
	a.Equal(tc, 3)
	pairs := []*file.Pair{}
	for _, pss := range pMap {
		pairs = append(pairs, pss...)
	}
	m.Add(pairs...)
	a.Equal(m.Count(), len(pairs))
}
