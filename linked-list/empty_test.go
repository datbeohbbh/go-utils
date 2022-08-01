package linkedlist

import "testing"

func TestEmpty(t *testing.T) {
	li := New[Integer32]()

	if !li.empty() {
		t.Error("linked list should be empty")
	}

	li.Insert(NewInteger(10))
	if li.empty() {
		t.Error("linked list should not be empty")
	}

	li.Remove(NewInteger(10))
	if !li.empty() {
		t.Error("linked list should be empty")
	}
}
