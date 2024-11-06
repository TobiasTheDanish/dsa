package lists

import (
	"fmt"
	"iter"
)

type ArrayList[T any] struct {
	data []T
	cap  int
}

var _ List[int] = (*ArrayList[int])(nil)

func NewArrayList[T any]() *ArrayList[T] {
	cap := 8

	return &ArrayList[T]{
		data: make([]T, 0, cap),
		cap:  cap,
	}
}

func (l *ArrayList[T]) Get(index int) (res T, err error) {
	if index < 0 || index >= len(l.data) {
		err = IndexOutOfBounds
		return
	}

	res = l.data[index]

	return res, err
}
func (l *ArrayList[T]) Add(e T) error {
	if len(l.data) >= l.cap {
		l.cap *= 2
		nData := make([]T, len(l.data), l.cap)
		copy(nData, l.data)
		l.data = nData
	}

	l.data = append(l.data, e)

	return nil
}

func (l *ArrayList[T]) Remove(index int) (res T, err error) {
	if index < 0 || index >= len(l.data) {
		err = IndexOutOfBounds
		return
	}

	res = l.data[index]

	for i := range len(l.data) - index {
		if index+i < len(l.data)-1 {
			l.data[index+i] = l.data[index+i+1]
		}
	}

	l.data = l.data[:len(l.data)-1]

	return res, err
}

func (l *ArrayList[T]) Size() int {
	return len(l.data)
}

func (l *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", l.data)
}

func (l *ArrayList[T]) Items() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, e := range l.data {
			if !yield(i, e) {
				return
			}
		}
	}
}
