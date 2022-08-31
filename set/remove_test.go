package avltree

import (
	"testing"

	"github.com/datbeohbbh/go-utils/numbers"
)

func TestRemove(t *testing.T) {
	set := New[T]()
	n := 1000000
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger(int32(i))
		set.Insert(v)
	}

	for i := n; i >= 1; i-- {
		v := numbers.NewInteger(int32(i))
		set.Insert(v)
	}

	removeCnt := 10000
	val := 0
	for i := 1; i <= removeCnt; i++ {
		v := numbers.NewInteger[int32](int32(val + 10))
		set.Remove(v)
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
		v := numbers.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		set.Remove(v)
		if set.Size() != int64(n-i) {
			t.Errorf("expected: %d, found: %d", n-i, set.Size())
		}
	}
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		ok := set.Remove(v)
		if ok {
			t.Error("expected: empty set")
		}
	}
}
