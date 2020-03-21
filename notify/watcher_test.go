package notify_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/meian/gowatch/notify"
	"github.com/meian/gowatch/testutil"
	"github.com/stretchr/testify/assert"
)

var e = notify.Export

func TestNewWatcher(t *testing.T) {
	a := assert.New(t)
	w, err := notify.NewWatcher()
	a.NoError(err)
	a.NotNil(w)
}

func TestNewWatcherError(t *testing.T) {
	a := assert.New(t)
	defer e.MockNewWatcherError()()
	w, err := notify.NewWatcher()
	a.Error(err)
	a.Nil(w)
}

func TestWatcherAdd(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			setupAlreadyAdded(w)
			defer w.Close()
			p := testDir(tt.path)
			err = w.Add(p)
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

func TestWatcherAddClosed(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			setupAlreadyAdded(w)
			w.Close()
			p := testDir(tt.path)
			err = w.Add(p)
			a.Error(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherAddError(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			defer w.Close()
			e.MockAddError(w)
			p := testDir(tt.path)
			err = w.Add(p)
			a.Error(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherRemove(t *testing.T) {
	// 監視削除は監視対象でも非対象でもエラーにならない
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			setupAlreadyAdded(w)
			defer w.Close()
			p := testDir(tt.path)
			err = w.Remove(p)
			a.NoError(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherRemoveClosed(t *testing.T) {
	// 監視削除はClose後もエラーにならない
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			setupAlreadyAdded(w)
			w.Close()
			p := testDir(tt.path)
			err = w.Remove(p)
			a.NoError(err)
			a.False(w.Watched(p))
		})
	}
}

func TestWatcherRemoveError(t *testing.T) {
	testutil.ChCurrentDir()
	tests := newTestData()
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if !tt.canAdd {
				// 監視が設定されないケースではエラーケースのテストはできない
				t.SkipNow()
			}
			a := assert.New(t)
			w, err := notify.NewWatcher()
			a.NoError(err)
			defer w.Close()
			e.MockRemoveError(w)
			p := testDir(tt.path)
			w.Add(p)
			err = w.Remove(p)
			a.Error(err)
			a.True(w.Watched(p))
		})
	}
}

func setupAlreadyAdded(w *notify.Watcher) {
	w.Add(testDir("already_added"))
}

func testDir(path string) string {
	cur, _ := os.Getwd()
	cur = filepath.Dir(cur)
	return filepath.Join(cur, "testdata", "watchtest", path)
}

type testData struct {
	desc   string
	path   string
	canAdd bool
}

func newTestData() []testData {
	return []testData{
		{desc: "exists", path: "exists", canAdd: true},
		{desc: "not exists", path: "not_exists", canAdd: false},
		{desc: "already added", path: "already_added", canAdd: true},
	}
}
