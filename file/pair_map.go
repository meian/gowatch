package file

import "sync"

// PairMap はPairを保持するマップ
type PairMap struct {
	m   map[string]*Pair
	mtx *sync.Mutex
}

// NewPairMap はPairMapを作成する
func NewPairMap() *PairMap {
	return &PairMap{
		m:   map[string]*Pair{},
		mtx: &sync.Mutex{},
	}
}

// Add はPairをマップに追加する
func (m *PairMap) Add(slice ...*Pair) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, p := range slice {
		if _, ok := m.m[p.Detected]; !ok {
			m.m[p.Detected] = p
		}
	}
}

// Has はファイル名のペアがマップに含まれるかを返す
func (m *PairMap) Has(detected string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	_, ok := m.m[detected]
	return ok
}

// Count はマップ内の要素数を返す
func (m *PairMap) Count() int {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	return len(m.m)
}

// PopAll はマップ内のペアをスライスで返し、マップ内の要素をクリアする
func (m *PairMap) PopAll() (pMap map[string][]*Pair, noTests []string) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	pMap = map[string][]*Pair{}
	noTests = []string{}
	for _, p := range m.m {
		if !p.TestEnabled() {
			noTests = append(noTests, p.Detected)
		} else {
			if _, ok := pMap[p.Test]; !ok {
				pMap[p.Test] = []*Pair{}
			}
			pMap[p.Test] = append(pMap[p.Test], p)
		}
	}
	m.m = map[string]*Pair{}
	return
}
