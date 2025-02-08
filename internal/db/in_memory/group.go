package in_memory

type GroupStore[K comparable, V any] struct {
	store *Store[K, V]
}

func NewGroupStore[K comparable, V any]() GroupStore[K, V] {
	return GroupStore[K, V]{store: NewStore[K, V]()}
}
