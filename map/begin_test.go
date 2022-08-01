package treap

import (
	"math/rand"
	"testing"
	"time"
)

func TestBegin(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[T, int]()
	n := 100000
	l := 1000 + rand.Intn(10000)
	r := 1000000000

	for i := 1; i <= l; i++ {
		treeMap.Put(T{value: i}, i)
	}

	for i := 1; i <= n; i++ {
		v := l + 1 + rand.Intn(r-l+1)
		treeMap.Put(T{value: v}, v)
	}

	for i := 1; i <= l; i++ {
		key, value := treeMap.Begin()
		if (key == nil || value == nil) || (*key != T{value: i} || *value != i) {
			t.Errorf("expected: {key: %d, value: %d}, found: {key: %v, value: %v}", i, i, key, value)
		}
		treeMap.Erase(*key)
	}

	treeMap.Clear()
	key, value := treeMap.Begin()
	if key != nil || value != nil {
		t.Error("{key, value} pair should be <nil>")
	}
}
