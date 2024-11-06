package heaps

import (
	"math"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewMinHeap[float64]()

	heap.Insert(42.0, 42)
	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	thirtyNine, _ := heap.Find()
	if thirtyNine != 39.0 {
		t.Fatalf("Elements not inserted correctly in min heap")
	}
	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio > e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	heap.Insert(15.0, 15)
	heap.Insert(42.0, 42)
	fifteen, _ := heap.Find()
	if fifteen != 15.0 {
		t.Fatalf("Elements not inserted correctly in min heap")
	}
	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio > e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	heap.Insert(15.0, 15)

	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio > e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}

	heap.Delete()

	for range heap.Size() {
		for i, e := range heap.elements {
			if i == 0 {
				continue
			}

			parentIndex := calcParentIndex(i)
			parent := heap.elements[parentIndex]

			if parent.prio > e.prio {
				t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
			}
		}
		heap.Delete()
	}

	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	heap.Insert(15.0, 15)

	fifteen, _ = heap.Replace(42.0, 42)
	if fifteen != 15.0 {
		t.Fatalf("Elements not inserted correctly in min heap")
	}

	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio > e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	v1, _ := heap.Extract()
	v2, _ := heap.Extract()
	v3, _ := heap.Extract()
	if v1 != 39.0 || v2 != 42.0 || v3 != 69.0 {
		t.Fatalf("Elements not extracted correctly from min heap")
	}
}

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap[float64]()

	heap.Insert(42.0, 42)
	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	sixtyNine, _ := heap.Find()
	if sixtyNine != 69.0 {
		t.Fatalf("Elements not inserted correctly in max heap")
	}
	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio < e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	heap.Insert(70.0, 70)
	heap.Insert(42.0, 42)
	seventy, _ := heap.Find()
	if seventy != 70.0 {
		t.Fatalf("Elements not inserted correctly in max heap")
	}
	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio < e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	heap.Insert(15.0, 15)

	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio < e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}

	heap.Delete()

	for range heap.Size() {
		for i, e := range heap.elements {
			if i == 0 {
				continue
			}

			parentIndex := calcParentIndex(i)
			parent := heap.elements[parentIndex]

			if parent.prio < e.prio {
				t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
			}
		}
		heap.Delete()
	}

	heap.Insert(69.0, 69)
	heap.Insert(39.0, 39)
	heap.Insert(70.0, 70)

	seventy, _ = heap.Replace(42.0, 42)
	if seventy != 70.0 {
		t.Fatalf("Elements not inserted correctly in max heap")
	}

	for i, e := range heap.elements {
		if i == 0 {
			continue
		}

		parentIndex := calcParentIndex(i)
		parent := heap.elements[parentIndex]

		if parent.prio < e.prio {
			t.Fatalf("Element at %d had lower prio than parent at %d. elem: '%v', parent: '%v'", i, parentIndex, e, parent)
		}
	}
	v3, _ := heap.Extract()
	v2, _ := heap.Extract()
	v1, _ := heap.Extract()
	if v1 != 39.0 || v2 != 42.0 || v3 != 69.0 {
		t.Fatalf("Elements not extracted correctly from max heap")
	}
}

func calcParentIndex(nodeIndex int) int {
	return int(math.Floor(float64((nodeIndex - 1) / 2)))
}
