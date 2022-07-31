package avltree

import (
	"math/rand"
	"testing"
	"time"
)

func TestLowerBoundSmall(t *testing.T) {
	a := []int{10, 9, 8, 16, 13, 17}
	set := New[T]()
	for _, v := range a {
		set.Insert(T{value: v})
	}
	res, _ := set.LowerBound(T{value: 10})
	if res == nil || (*res).value != 10 {
		t.Errorf("expected: %d, found: %v", 10, res)
	}

	res, _ = set.LowerBound(T{value: 9})
	if res == nil || (*res).value != 9 {
		t.Errorf("expected: %d, found: %v", 9, res)
	}

	res, _ = set.LowerBound(T{value: 14})
	if res == nil || (*res).value != 16 {
		t.Errorf("expected: %d, found: %v", 16, res)
	}

	res, _ = set.LowerBound(T{value: 12})
	if res == nil || (*res).value != 13 {
		t.Errorf("expected: %d, found: %v", 13, res)
	}

	res, _ = set.LowerBound(T{value: 20})
	if res != nil {
		t.Errorf("expected nil, found: %v", res)
	}

	set.Insert(T{value: 20})
	res, _ = set.LowerBound(T{value: 20})
	if res == nil {
		t.Errorf("expected %d, found: %v", 20, res)
	}
}

func TestLowerBoundStress(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	set := New[T]()

	n := 1000000
	l := 1000
	r := 1000000000

	for i := 1; i <= n; i++ {
		val := l + rand.Intn(r-l+1)
		set.Insert(T{value: val})
	}

	for val := 1; val < l; val++ {
		res, found := set.LowerBound(T{value: r + val})
		if found || res != nil {
			t.Errorf("expected: value %d can not be found", r+val)
		}
	}

	for val := 1; val < l; val++ {
		res, found := set.LowerBound(T{value: val})
		if !found || res == nil {
			t.Errorf("expected: value %d can be found", val)
		}
	}
}
