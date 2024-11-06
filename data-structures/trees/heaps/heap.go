package heaps

type Heap[T any] interface {
	Insert(T, int) error
	Find() (T, error)
	Extract() (T, error)
	Replace(T, int) (T, error)
	Delete() error
	Size() int
	IsEmpty() bool
	String() string
}
