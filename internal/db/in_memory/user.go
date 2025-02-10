package in_memory

type UserStore[K comparable, V any] struct {
	*Store[K, V]
}

func NewUserStore[K comparable, V any]() UserStore[K, V] {
	return UserStore[K, V]{NewStore[K, V]()}
}
