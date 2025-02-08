package db

type DB[K comparable, V any] interface {
	Create(K, V) error
	Update(K, V) error
	Delete(K) error
	Get(K) (V, error)
	List() (DB[K, V], error)
}
