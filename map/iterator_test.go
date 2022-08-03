package treap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/datbeohbbh/go-utils/types"
)

func TestIterator(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[types.Integer32, int64]()
	n := 100000
	l := 10000
	r := rand.Intn(1000000000)
	a := []int{}
	mp := make(map[int]int64)

	for i := 1; i <= n; i++ {
		key := l + rand.Intn(r-l+1)
		value := 1 + rand.Int63()
		if mp[key] == 0 {
			a = append(a, key)
		}
		mp[key] = value
		treeMap.Put(types.NewInteger32(key), value)
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	for repeat := 1; repeat <= 10; repeat++ {
		for next, c := treeMap.Iterator(), 0; c < len(a); c++ {
			key, value := next()
			if key == nil && value == nil {
				t.Error("expected: {key, value} != {nil, nil}")
			} else if key.GetValue() != int32(a[c]) && *value != mp[int(key.GetValue())] {
				t.Errorf("expected: {key, value}: {%d, %v}, found: {key, value}: {%d, %v}", a[c], mp[a[c]], key.GetValue(), *value)
			}
		}
	}
}
