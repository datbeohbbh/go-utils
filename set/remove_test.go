package avltree

import (
	"testing"
)

func TestRemove(t *testing.T) {
	set := New[T]()
	n := 1000000
	for i := 1; i <= n; i++ {
		set.Insert(T{value: i})
	}

	for i := n; i >= 1; i-- {
		set.Insert(T{value: i})
	}

	removeCnt := 10000
	val := 0
	for i := 1; i <= removeCnt; i++ {
		set.Remove(T{value: val + 10})
		val += 10
	}
	if set.Size() != int64(n-removeCnt) {
		t.Errorf("expected: %d, found: %d", n-removeCnt, set.Size())
	}
}

func TestRemoveEmptySet(t *testing.T) {
	set := New[T]()
	n := 1000
	for i := 1; i <= n; i++ {
		set.Insert(T{value: i})
	}
	for i := 1; i <= n; i++ {
		set.Remove(T{value: i})
		if set.Size() != int64(n-i) {
			t.Errorf("expected: %d, found: %d", n-i, set.Size())
		}
	}
	for i := 1; i <= n; i++ {
		ok := set.Remove(T{value: i})
		if ok {
			t.Error("expected: empty set")
		}
	}
}
