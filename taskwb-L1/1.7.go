package main

import "sync"

type Storage struct {
	data map[string]int
	mx   sync.RWMutex
}

func New() *Storage {
	return &Storage{
		data: make(map[string]int),
	}
}

func (s *Storage) Insert(key string, val int) {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.data[key] = val
}

func (s *Storage) Get(key string) (int, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	val, ok := s.data[key]
	return val, ok
}
