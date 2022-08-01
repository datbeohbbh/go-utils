package treap

import (
	"math/rand"
	"testing"
	"time"
)

func TestLowerBound(t *testing.T) {
	treeMap := New[T, int]()

	n := 100000
	for i := 2; i <= n; i += 2 {
		treeMap.Put(T{value: i}, i)
	}

	for i := 2; i <= n; i += 2 {
		key, value := treeMap.LowerBound(T{value: i})
		if key == nil || value == nil || (*key != T{value: i}) || *value != i {
			t.Errorf("expected: {key: %d, value: %d}, found: {key: %v, value: %v}", i, i, key, value)
		}

		key, value = treeMap.LowerBound(T{value: i - 1})
		if key == nil || value == nil || (*key != T{value: i}) || *value != i {
			t.Errorf("expected: {key: %d, value: %d}, found: {key: %v, value: %v}", i, i, key, value)
		}
	}
}

func TestLowerBoundFixedValue(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[T, int]()

	n := 100000
	l := 1
	r := l + rand.Intn(1000000000)

	for i := 1; i <= n; i++ {
		v := l + rand.Intn(r-l+1)
		treeMap.Put(T{value: v}, v)
	}

	key, _ := treeMap.LowerBound(T{value: r + 1})
	if key != nil {
		t.Errorf("expected: nil, found: %v", key)
	}
}
