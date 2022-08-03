package linkedlist

import "testing"

func TestEmpty(t *testing.T) {
	li := New[Integer32]()

	if !li.Empty() {
		t.Error("linked list should be Empty")
	}

	li.Insert(NewInteger(10))
	if li.Empty() {
		t.Error("linked list should not be Empty")
	}

	li.Remove(NewInteger(10))
	if !li.Empty() {
		t.Error("linked list should be Empty")
	}
}
