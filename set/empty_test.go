package avltree

import "testing"

func TestEmpty(t *testing.T) {
	set := New[T]()
	if !set.Empty() {
		t.Error("set must be empty")
	}

	n := 1000
	for i := 1; i <= n; i++ {
		set.Insert(T{value: i})
	}
	if set.Empty() {
		t.Error("set must not be empty")
	}

	for i := 1; i <= n+n; i++ {
		set.Remove(T{value: i})
	}
	if !set.Empty() {
		t.Error("set must be empty")
	}

	for i := 1; i <= n; i++ {
		set.Insert(T{value: i})
	}
	set.Clear()
	if !set.Empty() {
		t.Error("set must be empty")
	}
}
