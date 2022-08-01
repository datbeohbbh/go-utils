package avltree

import (
	"testing"

	"github.com/datbeohbbh/go-utils/types"
)

func TestUpperBoundSmall(t *testing.T) {
	set := New[T]()
	a := []int{4, 5, 6, 8, 9, 10, 13, 14, 15, 16, 17, 18, 20, 21}
	for _, v := range a {
		vv := types.NewInteger[int32](int32(v))
		set.Insert(vv)
	}

	res, _ := set.UpperBound(types.NewInteger[int32](int32(9)))
	if res == nil || (*res).GetValue() != 10 {
		t.Errorf("expected: %d, found %v", 10, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(3)))
	if res == nil || (*res).GetValue() != 4 {
		t.Errorf("expected: %d, found %v", 4, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(7)))
	if res == nil || (*res).GetValue() != 8 {
		t.Errorf("expected: %d, found %v", 8, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(12)))
	if res == nil || (*res).GetValue() != 13 {
		t.Errorf("expected: %d, found %v", 13, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(19)))
	if res == nil || (*res).GetValue() != 20 {
		t.Errorf("expected: %d, found %v", 20, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(13)))
	if res == nil || (*res).GetValue() != 14 {
		t.Errorf("expected: %d, found %v", 14, res)
	}

	res, _ = set.UpperBound(types.NewInteger[int32](int32(23)))
	if res != nil {
		t.Errorf("expected: nil, found %v", res)
	}
}

func TestUpperBoundStress(t *testing.T) {
	set := New[T]()
	l := 1
	r := 1000000
	for i := l; i <= r; i += 2 {
		v := types.NewInteger[int32](int32(i))
		set.Insert(v)
	}

	for i := l - 1; i < r; i += 2 {
		v := types.NewInteger[int32](int32(i))
		res, found := set.UpperBound(v)
		if found == false || (*res).GetValue() != int32(i+1) {
			t.Errorf("expected: %d, found: %v", i+1, res)
		}
	}
}

func TestUpperBoundPair(t *testing.T) {
	set := New[Pair]()
	set.Insert(Pair{
		first:  1,
		second: 2,
	})
	set.Insert(Pair{
		first:  2,
		second: 3,
	})

	_, found := set.UpperBound(Pair{first: 1, second: 0})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should be found", 1, 0)
	}

	_, found = set.UpperBound(Pair{first: 1, second: 1})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should be found", 1, 1)
	}

	_, found = set.UpperBound(Pair{first: 2, second: 4})
	if found {
		t.Errorf("With pair: {first: %v, second: %v} should not be found", 2, 4)
	}

	_, found = set.UpperBound(Pair{first: 2, second: 2})
	if !found {
		t.Errorf("With pair: {first: %v, second: %v} should not be found", 2, 2)
	}
}
