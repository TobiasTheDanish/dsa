package lists

import "iter"

type ArrayList[T any] struct {
	data []T
	cap  int
}

func NewArrayList[T any]() List[T] {
	cap := 8

	return &ArrayList[T]{
		data: make([]T, 0, cap),
		cap:  cap,
	}
}

func (l *ArrayList[T]) Get(int) (T, error)
func (l *ArrayList[T]) Add(T) error
func (l *ArrayList[T]) Remove(int) (T, error)
func (l *ArrayList[T]) Size() int
func (l *ArrayList[T]) Items() iter.Seq2[int, T]
