package sugar

import "sync"

// SMap goroutine safety map
type SMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func NewSMap[K comparable, V any]() *SMap[K, V] {
	return &SMap[K, V]{
		data: make(map[K]V),
	}
}

func (s *SMap[K, V]) Set(k K, v V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = v
}

func (s *SMap[K, V]) Get(k K) (V, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, exist := s.data[k]
	return v, exist
}

func (s *SMap[K, V]) Delete(k K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, k)
}

func (s *SMap[K, V]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *SMap[K, V]) ForEach(f func(K, V)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for k, v := range s.data {
		f(k, v)
	}
}
