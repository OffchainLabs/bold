package threadsafe

import "sync"

type Map[K comparable, V any] struct {
	sync.RWMutex
	items map[K]V
}

func NewMap[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{items: make(map[K]V)}
}

func (s *Map[K, V]) Keys() []K {
	s.RLock()
	defer s.RUnlock()
	keys := make([]K, 0, len(s.items))
	for key := range s.items {
		keys = append(keys, key)
	}
	return keys
}

func (s *Map[K, V]) Insert(k K, v V) {
	s.Lock()
	defer s.Unlock()
	s.items[k] = v
}

func (s *Map[K, V]) Get(k K) (V, bool) {
	s.RLock()
	defer s.RUnlock()
	item, ok := s.items[k]
	return item, ok
}

func (s *Map[K, V]) Has(k K) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.items[k]
	return ok
}

func (s *Map[K, V]) Delete(k K, v V) {
	s.Lock()
	defer s.Unlock()
	delete(s.items, k)
}

func (s *Map[K, V]) CopyItems() map[K]V {
	s.RLock()
	defer s.RUnlock()
	copied := make(map[K]V, len(s.items))
	for k, v := range s.items {
		copied[k] = v
	}
	return copied
}

type Set[T comparable] struct {
	sync.RWMutex
	items map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]bool),
	}
}

func (s *Set[T]) Insert(t T) {
	s.Lock()
	defer s.Unlock()
	s.items[t] = true
}

func (s *Set[T]) Has(t T) bool {
	s.RLock()
	defer s.RUnlock()
	return s.items[t]
}

func (s *Set[T]) CopyItems() map[T]bool {
	s.RLock()
	defer s.RUnlock()
	copied := make(map[T]bool, len(s.items))
	for k, v := range s.items {
		copied[k] = v
	}
	return copied
}
