package avltree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/datbeohbbh/go-utils/numbers"
)

func TestLowerBoundSmall(t *testing.T) {
	a := []int{10, 9, 8, 16, 13, 17}
	set := New[T]()
	for _, v := range a {
		vv := numbers.NewInteger(int32(v))
		set.Insert(vv)
	}
	res, _ := set.LowerBound(numbers.NewInteger(int32(10)))
	if res == nil || (*res).GetValue() != 10 {
		t.Errorf("expected: %d, found: %v", 10, res)
	}

	res, _ = set.LowerBound(numbers.NewInteger(int32(9)))
	if res == nil || (*res).GetValue() != 9 {
		t.Errorf("expected: %d, found: %v", 9, res)
	}

	res, _ = set.LowerBound(numbers.NewInteger(int32(14)))
	if res == nil || (*res).GetValue() != 16 {
		t.Errorf("expected: %d, found: %v", 16, res)
	}

	res, _ = set.LowerBound(numbers.NewInteger(int32(12)))
	if res == nil || (*res).GetValue() != 13 {
		t.Errorf("expected: %d, found: %v", 13, res)
	}

	res, _ = set.LowerBound(numbers.NewInteger(int32(20)))
	if res != nil {
		t.Errorf("expected nil, found: %v", res)
	}

	set.Insert(numbers.NewInteger(int32(20)))
	res, _ = set.LowerBound(numbers.NewInteger(int32(20)))
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
		v := numbers.NewInteger(int32(val))
		set.Insert(v)
	}

	for val := 1; val < l; val++ {
		v := numbers.NewInteger(int32(r + val))
		res, found := set.LowerBound(v)
		if found || res != nil {
			t.Errorf("expected: GetValue() %d can not be found", r+val)
		}
	}

	for val := 1; val < l; val++ {
		v := numbers.NewInteger(int32(val))
		res, found := set.LowerBound(v)
		if !found || res == nil {
			t.Errorf("expected: GetValue() %d can be found", val)
		}
	}
}
