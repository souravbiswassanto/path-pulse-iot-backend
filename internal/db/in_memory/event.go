package in_memory

type EventStore[K comparable, V any] struct {
	*Store[K, V]
}

func NewEventStore[K comparable, V any]() EventStore[K, V] {
	return EventStore[K, V]{NewStore[K, V]()}
}
