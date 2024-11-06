package lists

import "testing"

func TestArrayList(t *testing.T) {
	list := NewArrayList[uint]()

	list.Add(42)
	list.Add(69)

	if list.Size() != 2 {
		t.Fatalf("Elements were not added properly. Expected size of 2, but got %d", list.Size())
	}

	head, err := list.Get(0)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if head != 42 {
		t.Fatalf("Unexpected value at index 0")
	}

	list.Add(420)

	sixtyNine, err := list.Remove(1)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if sixtyNine != 69 {
		t.Fatalf("Unexpected value after remove")
	}
	if list.Size() != 2 {
		t.Fatalf("Elements were not removed properly. Expected size of 2, but got %d", list.Size())
	}

	tail, err := list.Get(1)
	if err != nil {
		t.Fatalf("Got unexpected error: %v", err)
	}
	if tail != 420 {
		t.Fatalf("Unexpected value at index 1 after remove")
	}

	for i, num := range list.Items() {
		if i == 0 && num != 42 {
			t.Fatalf("Unexpected value '%d' at index '%d'. '%s'", num, i, list.String())
		}
		if i == 1 && num != 420 {
			t.Fatalf("Unexpected value '%d' at index '%d'. '%s'", num, i, list.String())
		}
	}
}
