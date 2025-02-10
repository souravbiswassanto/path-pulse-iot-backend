package in_memory

type TrackerStore[K comparable, M comparable, V any] struct {
	Checkpoint *CheckpointStore[K, V]
	Position   *PositionStore[M, V]
}

func NewTrackerStore[K comparable, M comparable, V any]() TrackerStore[K, M, V] {
	cStore := &CheckpointStore[K, V]{
		NewStore[K, V](),
	}
	pStore := &PositionStore[M, V]{
		NewStore[M, V](),
	}
	return TrackerStore[K, M, V]{
		Checkpoint: cStore,
		Position:   pStore,
	}
}

type CheckpointStore[K comparable, V any] struct {
	*Store[K, V]
}

type PositionStore[K comparable, V any] struct {
	*Store[K, V]
}
