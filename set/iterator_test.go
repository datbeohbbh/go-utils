package avltree

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/datbeohbbh/go-utils/numbers"
)

func TestIterator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeSet := New[T]()
	n := 100000
	l := 1000
	r := l + rand.Intn(1000000000)
	a := []int{}
	has := make(map[int]bool)
	for i := 1; i <= n; i++ {
		rval := l + rand.Intn(r-l+1)
		treeSet.Insert(numbers.NewInteger32(rval))
		if !has[rval] {
			a = append(a, rval)
			has[rval] = true
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return (a[i] < a[j])
	})

	for repeat := 1; repeat <= 5; repeat++ {
		for next, c := treeSet.Iterator(), 0; c < len(a); c++ {
			val := next()
			if val == nil || val.GetValue() != int32(a[c]) {
				if val == nil {
					t.Error("expected: val is not nil")
				} else {
					t.Errorf("expected: sorted slice, found at position c = %d: a[c] = %d val = %d", c, a[c], val.GetValue())
				}
			}
		}
	}
}
