package avltree

import (
	"testing"

	"github.com/datbeohbbh/go-utils/types"
)

func TestInsertSmall(t *testing.T) {
	set := New[T]()
	for i := 1; i <= 10; i++ {
		v := types.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	if set.Size() != 10 {
		t.Errorf("expected: 10, found %d", set.Size())
	}
	for i := 1; i <= 20; i++ {
		v := types.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	if set.Size() != 20 {
		t.Errorf("expected: 20, found %d", set.Size())
	}
}

func TestInsertWithDuplicateValye(t *testing.T) {
	set := New[T]()
	repeat := 10
	for ; repeat > 0; repeat-- {
		for i := 1; i <= 50; i++ {
			v := types.NewInteger[int32](int32(i))
			set.Insert(v)
		}
	}
	if set.Size() != 50 {
		t.Errorf("expected: 50, found %d", set.Size())
	}
}

func TestInsertPair(t *testing.T) {
	set := New[Pair]()
	n := 10
	expected := n * n
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			set.Insert(Pair{
				first:  i,
				second: j,
			})
		}
	}
	inserted := set.Insert(Pair{
		first:  1,
		second: 1,
	})
	if inserted {
		t.Errorf("expected: insert failed")
	}
	if set.Size() != int64(expected) {
		t.Errorf("expected: %v, found %v", expected, set.Size())
	}
}

func TestInsertStress(t *testing.T) {
	set := New[T]()
	n := 1000000
	expectedSize := 1000000 + 10
	for i := 1; i <= n; i++ {
		v := types.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	for i := n + 1; i <= n+10; i++ {
		v := types.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	if set.Size() != int64(expectedSize) {
		t.Errorf("expected: %v, found: %v", expectedSize, set.Size())
	}
}
