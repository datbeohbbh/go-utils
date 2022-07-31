package avltree

import (
	"math/rand"
	"testing"
	"time"
)

func TestFindSmall(t *testing.T) {
	set := New[T]()
	n := 10
	for i := 1; i <= n; i++ {
		set.Insert(T{value: i})
	}
	for i := 1; i <= n; i++ {
		ok := set.Find(T{value: i})
		if !ok {
			t.Errorf("expected: value %d should be in set, found: %v", i, ok)
		}
	}
	for i := 1; i <= n; i++ {
		ok := set.Find(T{value: i + n})
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
		set.Insert(T{value: value})
		if i <= 1000 {
			a = append(a, value)
		}
	}

	for _, v := range a {
		ok := set.Find(T{value: v})
		if !ok {
			t.Errorf("expected: value %d should be in set, found: %v", v, ok)
		}
	}

	for _, v := range a {
		set.Remove(T{value: v})
	}

	for _, v := range a {
		ok := set.Find(T{value: v})
		if ok {
			t.Errorf("expected: value %d should not be in set, found: %v", v, ok)
		}
	}
}
