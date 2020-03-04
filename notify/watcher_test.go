package notify_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/meian/gowatch/notify"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

func TestWatcherAdd(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		w := newWatcher()
		defer w.Close()
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			p := testDir(tt.path)
			err := w.Add(p)
			if tt.canAdd {
				a.NoError(err, p)
				a.True(w.Watched(p))
				a.NoError(w.Remove(p))
				a.False(w.Watched(p))
			} else {
				a.Error(err, p)
				a.False(w.Watched(p))
			}
		})
	}
}

func TestWatcherRemove(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		w := newWatcher()
		defer w.Close()
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			p := testDir(tt.path)
			err := w.Remove(p)
			a.NoError(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherAddClosed(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		w := newWatcher()
		w.Close()
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			p := testDir(tt.path)
			err := w.Add(p)
			a.Error(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherRemoveClosed(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		w := newWatcher()
		w.Close()
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			p := testDir(tt.path)
			err := w.Remove(p)
			a.NoError(err)
			a.False(w.Watched(p))
		})
	}
}

func newWatcher() *notify.Watcher {
	w, err := notify.NewWatcher()
	if err != nil {
		panic(err)
	}
	w.Add(testDir("already_added"))
	return w
}

func testDir(path string) string {
	cur, _ := os.Getwd()
	cur = filepath.Dir(cur)
	return filepath.Join(cur, "internal", "watchtest", path)
}

type testData struct {
	desc      string
	path      string
	canAdd    bool
	canRemove bool
}

func newTestData() []testData {
	return []testData{
		{desc: "exists", path: "exists", canAdd: true},
		{desc: "not exists", path: "not_exists", canAdd: false},
		{desc: "already added", path: "already_added", canAdd: true},
	}
}
