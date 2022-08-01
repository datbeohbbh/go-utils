package treap

import "testing"

func TestEmpty(t *testing.T) {
	treeMap := New[T, int]()
	if !treeMap.Empty() {
		t.Error("map should be empty")
	}

	treeMap.Put(T{value: 1}, 1)
	if treeMap.Empty() {
		t.Error("map should not be empty")
	}

	treeMap.Put(T{value: 1}, 1)
	treeMap.Erase(T{value: 2})
	treeMap.Erase(T{value: 1})
	if !treeMap.Empty() {
		t.Error("map should be empty")
	}
}
