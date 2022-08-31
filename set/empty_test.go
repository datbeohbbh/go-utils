package avltree

import (
	"testing"

	"github.com/datbeohbbh/go-utils/numbers"
)

func TestEmpty(t *testing.T) {
	set := New[T]()
	if !set.Empty() {
		t.Error("set must be empty")
	}

	n := 1000
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	if set.Empty() {
		t.Error("set must not be empty")
	}

	for i := 1; i <= n+n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		set.Remove(v)
	}
	if !set.Empty() {
		t.Error("set must be empty")
	}

	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	set.Clear()
	if !set.Empty() {
		t.Error("set must be empty")
	}
}
