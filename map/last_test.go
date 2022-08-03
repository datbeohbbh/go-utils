package treap

import "testing"

func TestLast(t *testing.T) {
	treeMap := New[T, int]()

	n := 1000000
	for i := n; i >= 1; i-- {
		treeMap.Put(T{value: i}, i)
		key, value := treeMap.Last()
		if key == nil || value == nil || (*key != T{value: n}) || *value != n {
			t.Errorf("expected: {key: %d, value: %d}, found {key: %v, value: %v}", n, n, key, value)
		}
	}

	for i := n; i >= 2; i-- {
		treeMap.Erase(T{value: i})
		key, value := treeMap.Last()
		if key == nil || value == nil || (*key != T{value: i - 1}) || *value != i-1 {
			t.Errorf("expected: {key: %d, value: %d}, found {key: %v, value: %v}", n, n, key, value)
		}
	}
}
