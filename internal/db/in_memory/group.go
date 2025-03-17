package in_memory

type GroupStore[K comparable, V any] struct {
	*Store[K, V]
}

func NewGroupStore[K comparable, V any](s *Store[K, V]) GroupStore[K, V] {
	return GroupStore[K, V]{s}
}
