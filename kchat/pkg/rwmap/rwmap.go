package rwmap

import (
	"sync"
)

type RWMap struct {
	sync.RWMutex
	m map[int]interface{}
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]interface{}, n),
	}
}

func (m *RWMap) Get(k int) (interface{}, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap) Set(k int, v interface{}) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RWMap) Delete(k int) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RWMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RWMap) Each(f func(k int, v interface{}) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
