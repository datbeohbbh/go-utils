package avltree

import (
	"math/rand"
	"testing"
	"time"

	"github.com/datbeohbbh/go-utils/numbers"
)

func TestFindSmall(t *testing.T) {
	set := New[T]()
	n := 10
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		set.Insert(v)
	}
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i))
		ok := set.Find(v)
		if !ok {
			t.Errorf("expected: value %d should be in set, found: %v", i, ok)
		}
	}
	for i := 1; i <= n; i++ {
		v := numbers.NewInteger[int32](int32(i + n))
		ok := set.Find(v)
		if ok {
			t.Errorf("expected: value %d should not be in set, found: %v", i, ok)
		}
	}
}

func TestFindStress(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	set := New[T]()
	n := 1000000
	l := 1
	r := 1000000000
	a := []int{}

	for i := 1; i <= n; i++ {
		value := l + rand.Intn(r-l+1)
		v := numbers.NewInteger[int32](int32(value))
		set.Insert(v)
		if i <= 1000 {
			a = append(a, value)
		}
	}

	for _, v := range a {
		vv := numbers.NewInteger[int32](int32(v))
		ok := set.Find(vv)
		if !ok {
			t.Errorf("expected: value %d should be in set, found: %v", v, ok)
		}
	}

	for _, v := range a {
		vv := numbers.NewInteger[int32](int32(v))
		set.Remove(vv)
	}

	for _, v := range a {
		vv := numbers.NewInteger[int32](int32(v))
		ok := set.Find(vv)
		if ok {
			t.Errorf("expected: value %d should not be in set, found: %v", v, ok)
		}
	}
}
