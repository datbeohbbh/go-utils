package treap

import "testing"

func TestGet(t *testing.T) {
	treeMap := New[T, int64]()

	n := 100000
	for i := 1; i <= n; i++ {
		treeMap.Put(T{value: i}, int64(i)*int64(i))
	}

	for i := 1; i <= n; i++ {
		v := treeMap.Get(T{value: i})
		if v == nil || *v != int64(i)*int64(i) {
			t.Errorf("expected: %v, found: %v", int64(i)*int64(i), v)
		}
	}

	for i := n/2 + 1; i <= n; i++ {
		treeMap.Put(T{value: i}, int64(i))
	}

	for i := 1; i <= n/2; i++ {
		v := treeMap.Get(T{value: i})
		if v == nil || *v != int64(i)*int64(i) {
			t.Errorf("expected: %v, found: %v", int64(i)*int64(i), v)
		}
	}

	for i := (n / 2) + 1; i <= n; i++ {
		v := treeMap.Get(T{value: i})
		if v == nil || *v != int64(i) {
			t.Errorf("expected: %v, found: %v", int64(i), v)
		}
	}
}
