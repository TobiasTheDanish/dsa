package heaps

import (
	"errors"
	"fmt"
	"math"
)

var _ Heap[float32] = (*MaxHeap[float32])(nil)

type MaxHeap[T any] struct {
	elements []heapNode[T]
}

func NewMaxHeap[T any]() *MaxHeap[T] {
	return &MaxHeap[T]{
		elements: make([]heapNode[T], 0),
	}
}

func (h *MaxHeap[T]) Insert(e T, prio int) error {
	node := heapNode[T]{
		prio: prio,
		data: e,
	}

	h.elements = append(h.elements, node)

	nodeIndex := len(h.elements) - 1
	if nodeIndex == 0 {
		// this is the root, return
		return nil
	}

	h.siftUp(nodeIndex)

	return nil
}

func (h *MaxHeap[T]) Find() (T, error) {
	return h.elements[0].data, nil
}

func (h *MaxHeap[T]) Extract() (res T, err error) {
	if len(h.elements) == 0 {
		err = errors.New("Cannot extract from empty heap")
		return
	}

	res = h.elements[0].data
	err = h.Delete()
	return
}

func (h *MaxHeap[T]) Replace(data T, prio int) (res T, err error) {
	if len(h.elements) == 0 {
		err = errors.New("Cannot replace in empty heap")
		return
	}

	res = h.elements[0].data
	h.elements[0].data = data
	h.elements[0].prio = prio

	h.siftDown(0)

	return
}

func (h *MaxHeap[T]) Delete() (err error) {
	if len(h.elements) == 0 {
		err = errors.New("Cannot delete from empty heap")
		return
	}

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	h.siftDown(0)

	return
}

func (h *MaxHeap[T]) Size() int {
	return len(h.elements)
}

func (h *MaxHeap[T]) IsEmpty() bool {
	return len(h.elements) > 0
}

func (h *MaxHeap[T]) String() string {
	return fmt.Sprintf("%v", h.elements)
}

func (h *MaxHeap[T]) siftUp(index int) {
	nodeIndex := index
	if nodeIndex == 0 {
		// this is the root, return
		return
	}

	parentIndex := int(math.Floor(float64((nodeIndex - 1) / 2)))
	parent := h.elements[parentIndex]
	for parent.prio < h.elements[nodeIndex].prio {
		// swap
		h.elements[parentIndex] = h.elements[nodeIndex]
		h.elements[nodeIndex] = parent
		nodeIndex = parentIndex
		if nodeIndex == 0 {
			// the new element has been sifted up to root
			break
		}
		parentIndex = int(math.Floor(float64((nodeIndex - 1) / 2)))

		parent = h.elements[parentIndex]
	}
}

func (h *MaxHeap[T]) siftDown(index int) {
	nodeIndex := index
	childIndex1 := nodeIndex*2 + 1
	childIndex2 := nodeIndex*2 + 2

	for childIndex1 < len(h.elements) && childIndex2 < len(h.elements) {
		node := h.elements[nodeIndex]
		child1 := h.elements[childIndex1]
		child2 := h.elements[childIndex2]

		if child1.prio <= node.prio && child2.prio <= node.prio {
			// the element has reached a fitting position
			break
		}

		if child1.prio > node.prio && child2.prio > node.prio {
			// choose highest child to swap with
			if child1.prio > child2.prio {
				h.elements[nodeIndex] = child1
				h.elements[childIndex1] = node
				nodeIndex = childIndex1
			} else {
				h.elements[nodeIndex] = child2
				h.elements[childIndex2] = node
				nodeIndex = childIndex2
			}
		} else if child1.prio > node.prio {
			//swap
			h.elements[nodeIndex] = child1
			h.elements[childIndex1] = node
			nodeIndex = childIndex1
		} else if child2.prio > node.prio {
			// swap
			h.elements[nodeIndex] = child2
			h.elements[childIndex2] = node
			nodeIndex = childIndex2
		}

		childIndex1 = nodeIndex*2 + 1
		childIndex2 = nodeIndex*2 + 2
	}

	if childIndex1 < len(h.elements) && childIndex2 >= len(h.elements) {
		node := h.elements[nodeIndex]
		child1 := h.elements[childIndex1]
		if child1.prio > node.prio {
			//swap
			h.elements[nodeIndex] = child1
			h.elements[childIndex1] = node
			nodeIndex = childIndex1
		}
	}
}
