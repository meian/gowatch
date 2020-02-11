package util

import "sync"

// StringSet はstringの重複なしのセットを表す
type StringSet struct {
	m   map[string]bool
	mtx *sync.Mutex
}

// NewStringSet はStringSetを作成する
func NewStringSet() *StringSet {
	ss := &StringSet{
		m:   map[string]bool{},
		mtx: &sync.Mutex{},
	}
	return ss
}

// Add は文字列を追加する
func (ss *StringSet) Add(s string) {
	ss.AddSlice([]string{s})
}

// AddSlice はスライスの文字列を追加する
func (ss *StringSet) AddSlice(list []string) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
	for _, s := range list {
		if _, ok := ss.m[s]; !ok {
			ss.m[s] = true
		}
	}
}

// Has は文字列がセットに含まれるかを判別する
func (ss *StringSet) Has(s string) bool {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
	_, ok := ss.m[s]
	return ok
}

// Count はセット内の要素数を返す
func (ss *StringSet) Count() int {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
	return len(ss.m)
}

// PopAll はセット内の要素をスライスで返し、セット内の要素をクリアする
func (ss *StringSet) PopAll() []string {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()
	a := make([]string, 0, len(ss.m))
	for s := range ss.m {
		a = append(a, s)
	}
	ss.m = map[string]bool{}
	return a
}
