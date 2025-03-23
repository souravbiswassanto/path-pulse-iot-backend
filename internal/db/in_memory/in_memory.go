package in_memory

import (
	"context"
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/pkg/lib"
	"sync"
	"time"
)

type Store[K comparable, V interface{}] struct {
	ctx   context.Context
	store map[K]V
	mu    sync.RWMutex // Here using RWMutex because Read >> Write
	*StoreOptions
}

type StoreOptions struct {
	MaxKeyStoreLimit int `yaml:"maxKeyStoreLimit,omitempty'"`
}

func NewStore[K comparable, V interface{}](ctx context.Context, opts *StoreOptions) *Store[K, V] {
	store := &Store[K, V]{
		ctx:          ctx,
		store:        make(map[K]V),
		StoreOptions: opts,
	}
	store.CleanOldCaches()
	return store
}

func (s *Store[K, V]) Create(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

// Get returns the value if exists,
// unless it returns a keyNotFound error
func (s *Store[K, V]) Get(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.store[key]
	if !ok {
		var zero V
		return zero, custom_error.ErrKeyNotFound
	}
	return v, nil
}

func (s *Store[K, V]) Update(key K, value V) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[key] = value
}

func (s *Store[K, V]) Delete(key K) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, key)
}

func (s *Store[K, V]) List() map[K]V {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m := make(map[K]V)
	for k, v := range s.store {
		m[k] = v
	}
	return m
}

func (s *Store[K, V]) CleanOldCaches() {
	ticker := time.NewTicker(time.Minute * 5)
	defer ticker.Stop()
	cutExtra := 500
	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			s.mu.RLock()
			size := len(s.store)
			s.mu.RUnlock()
			if size <= s.MaxKeyStoreLimit {
				continue
			}
			td := make([]K, 0)
			s.mu.RLock()
			for k, _ := range s.store {
				if size <= max(0, s.MaxKeyStoreLimit-cutExtra) {
					break
				}
				td = append(td, k)
				size--
			}
			s.mu.RUnlock()
			s.mu.Lock()
			for _, v := range td {
				delete(s.store, v)
			}
			s.mu.Unlock()
		}
	}
}
