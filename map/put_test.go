package treap

import (
	"math/rand"
	"testing"
	"time"
)

func TestPutSmall(t *testing.T) {
	treeMap := New[T, int]()
	n := 10

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			treeMap.Put(T{value: i}, j)
		}
	}

	if treeMap.Size() != int64(n) {
		t.Errorf("expected: %d, found: %d", n, treeMap.Size())
	}

	for i := 1; i <= n; i++ {
		val := treeMap.Get(T{value: i})
		if val == nil || *val != n {
			t.Errorf("expected: %d, found: %v", n, val)
		}
	}
}

func TestPutStress(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[T, int]()

	n := 1000000
	l := 100000
	r := 1000000000
	a := []int{}

	for i := 1; i <= n; i++ {
		val := l + rand.Intn(r-l+1)
		treeMap.Put(T{value: val}, i+rand.Intn(1000000000))
		if rand.Intn(2) == 1 {
			a = append(a, val)
		}
	}

	for _, v := range a {
		if !treeMap.Contains(T{value: v}) {
			t.Errorf("expected: true, found: false on key: %d", v)
		}
	}

	finalValue := make(map[int]int)
	for idx, v := range a {
		treeMap.Put(T{value: v}, idx)
		finalValue[v] = idx
	}

	for _, v := range a {
		val := treeMap.Get(T{value: v})
		if val == nil || *val != finalValue[v] {
			t.Errorf("expected: %d, found: %v", v, val)
		}
	}
}
