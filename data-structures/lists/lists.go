package lists

import (
	"errors"
	"iter"
)

var (
	IndexOutOfBounds = errors.New("Index out of bounds")
)

type List[T any] interface {
	// Get the element at index, returns error if index is out of bounds
	Get(int) (T, error)
	// Add an element to the end of the list
	Add(T) error
	// Remove the element at index, returns error if index is out of bounds
	Remove(int) (T, error)
	// Get the number of elements in the list, also known as the size of the list
	Size() int
	// Iterator function
	Items() iter.Seq2[int, T]
	String() string
}
