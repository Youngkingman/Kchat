package service

import "sync"

var ChatRoomMap = NewRWMap(0)

type RWMap struct {
	sync.RWMutex
	m map[int]*ChatRoom
}

func NewRWMap(n int) *RWMap {
	return &RWMap{
		m: make(map[int]*ChatRoom, n),
	}
}

func (m *RWMap) Get(k int) (*ChatRoom, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RWMap) Set(k int, v *ChatRoom) {
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

func (m *RWMap) Each(f func(k int, v *ChatRoom) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
