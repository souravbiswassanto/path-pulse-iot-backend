package in_memory

import (
	custom_error "github.com/souravbiswassanto/path-pulse-iot-backend/internal/custom-error"
	"sync"
)

type Store[K comparable, V interface{}] struct {
	store map[K]V
	mu    sync.RWMutex
}

func NewStore[K comparable, V interface{}]() *Store[K, V] {
	return &Store[K, V]{
		store: make(map[K]V),
	}
}

//type InMemoryStore struct {
//	eventStore      *Store[uint64, models.Event]
//	groupStore      *Store[uint64, models.Group]
//	userStore       *Store[uint64, models.User]
//	checkpointStore *Store[uint64, models.Checkpoint]
//	positionStore   *Store[*time.Time, models.Position]
//}
//
//func NewInMemoryStore() *InMemoryStore {
//	return &InMemoryStore{
//		eventStore:      NewStore[uint64, models.Event](),
//		groupStore:      NewStore[uint64, models.Group](),
//		userStore:       NewStore[uint64, models.User](),
//		checkpointStore: NewStore[uint64, models.Checkpoint](),
//		positionStore:   NewStore[*time.Time, models.Position](),
//	}
//}

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
