package linkedlist

import (
	"math/rand"
	"testing"
	"time"
)

func TestSize(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	li := New[Integer32]()
	n := 10000
	l := 100
	r := 100000
	a := []int{}
	has := make(map[int]bool)

	for i := 1; i <= n; i++ {
		val := l + rand.Intn(r-l+1)
		for has[val] {
			val = l + rand.Intn(r-l+1)
		}
		a = append(a, val)
		has[val] = true
		li.Insert(NewInteger(val))
		if li.Size() != int64(i) {
			t.Errorf("expected: %d, found: %d", i, li.Size())
		}
	}

	for len(a) > 0 {
		idx := rand.Intn(len(a))
		li.Remove(NewInteger(a[idx]))
		a = removeIndex(a, idx)
		if li.Size() != int64(len(a)) {
			t.Errorf("expected: %d, found: %d", len(a), li.Size())
		}
	}
}
