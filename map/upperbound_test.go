package treap

import (
	"math/rand"
	"testing"
	"time"
)

func TestUpperBound(t *testing.T) {
	treeMap := New[T, int]()

	n := 100000
	for i := 2; i <= n; i += 2 {
		treeMap.Put(T{value: i}, i)
	}

	for i := 1; i <= n; i += 2 {
		key, value := treeMap.UpperBound(T{value: i})
		if key == nil || value == nil || (*key != T{value: i + 1}) || *value != i+1 {
			t.Errorf("expected: {key: %d, value: %d}, found {key: %v, value: %v}", i+1, i+1, key, value)
		}
	}

	treeMap.Clear()
	key, value := treeMap.UpperBound(T{value: 1})
	if key != nil || value != nil {
		t.Error("expected: {key: nil, value: nil}")
	}
}

func TestUpperBoundWithErase(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[T, int]()

	n := 100000

	l := 1000 + rand.Intn(1000)

	for i := l + 1; i <= n; i++ {
		treeMap.Put(T{value: i}, i)
	}
	for i := l + 1; i <= n; i++ {
		v := 1 + rand.Intn(l)
		key, value := treeMap.UpperBound(T{value: v})
		if key == nil || value == nil || (*key != T{value: i}) || *value != i {
			t.Errorf("expected: {key: %d, value: %d}, found: {key: %v, value: %v}", i, i, key, value)
		}
		treeMap.Erase(*key)
	}
}
