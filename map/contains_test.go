package treap

import "testing"

func TestContains(t *testing.T) {
	treeMap := New[T, int]()

	n := 1000
	for i := 1; i <= n; i += 2 {
		treeMap.Put(T{value: i}, i)
	}

	for i := 2; i <= n; i += 2 {
		contains := treeMap.Contains(T{value: i})
		if contains {
			t.Error("expected: false, found: true")
		}
		contains = treeMap.Contains(T{value: i - 1})
		if !contains {
			t.Error("expected: true, found: false")
		}
	}
}
