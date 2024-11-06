package lists

import (
	"fmt"
	"iter"
	"strings"
)

var _ List[int] = (*LinkedList[int])(nil)

type node[T any] struct {
	data T
	next *node[T]
}

type LinkedList[T any] struct {
	size int
	head *node[T]
	tail *node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l *LinkedList[T]) Get(index int) (res T, err error) {
	if index < 0 || index >= l.size {
		err = IndexOutOfBounds
		return
	}

	if index == l.size-1 {
		res = l.tail.data
		return
	}

	n := l.head
	for range index {
		n = n.next
	}

	res = n.data

	return
}

func (l *LinkedList[T]) Add(data T) error {
	n := &node[T]{
		data: data,
	}

	if l.head == nil {
		l.head = n
	}

	if l.tail != nil {
		l.tail.next = n
	}

	l.tail = n

	l.size += 1

	return nil
}

func (l *LinkedList[T]) Remove(index int) (res T, err error) {
	if index < 0 || index >= l.size {
		err = IndexOutOfBounds
		return
	}

	if l.size == 1 {
		res = l.head.data
		l.head = nil
		l.tail = nil
		l.size -= 1
		return
	} else if index == 0 {
		res = l.head.data
		l.head = l.head.next
		l.size -= 1
		return
	}

	var prev *node[T]
	n := l.head
	for range index {
		prev = n
		n = n.next
	}

	prev.next = n.next
	n.next = nil
	res = n.data
	l.size -= 1

	return
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) String() string {
	sb := strings.Builder{}

	for _, e := range l.Items() {
		sb.WriteString(fmt.Sprintf("%v", e))
	}

	return sb.String()
}

func (l *LinkedList[T]) Items() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		n := l.head
		for i := range l.size {
			if !yield(i, n.data) {
				return
			}
			n = n.next
		}
	}
}
