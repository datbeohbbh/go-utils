package treap

import "testing"

func TestClear(t *testing.T) {
	treeMap := New[T, int]()

	n := 10000
	for i := 1; i <= n; i++ {
		treeMap.Put(T{value: i}, i)
	}

	treeMap.Clear()
	if !treeMap.Empty() {
		t.Error("map should be empty")
	}
}
