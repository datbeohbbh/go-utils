package treap

import (
	"math/rand"
	"testing"
	"time"
)

func TestErase(t *testing.T) {
	treeMap := New[T, int]()

	n := 100000
	for i := 1; i <= n; i++ {
		treeMap.Put(T{value: i}, i)
	}
	for i := 1; i <= (n / 2); i++ {
		ok := treeMap.Erase(T{value: i})
		if !ok {
			t.Error("expected: true, found: fail")
		}
	}
	if treeMap.Size() != int64(n/2) {
		t.Errorf("expected: %d, found: %d", n/2, treeMap.Size())
	}

	for rep := 1; rep <= 10; rep++ {
		for i := 1; i <= (n / 2); i++ {
			ok := treeMap.Erase(T{value: i})
			if ok {
				t.Error("expected: false, found: true")
			}
		}
	}

	for i := 1; i <= n; i++ {
		treeMap.Erase(T{value: i})
	}
	if !treeMap.Empty() {
		t.Error("expected: empty map, found: non-empty map")
	}
}

func TestEraseRandomly(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	treeMap := New[T, int]()

	n := 100000
	l := 1
	r := rand.Intn(1000000000)
	a := []int{}
	for i := 1; i <= n; i++ {
		v := l + rand.Intn(r-l+1)
		treeMap.Put(T{value: v}, v)
		if rand.Intn(100) > 50 {
			a = append(a, v)
		}
	}

	vis := make(map[int]bool)
	for _, v := range a {
		if !vis[v] {
			ok := treeMap.Erase(T{value: v})
			if !ok {
				t.Error("expected: true, found: false")
			}
			vis[v] = true
		} else {
			ok := treeMap.Erase(T{value: v})
			if ok {
				t.Error("expected: false, found: true")
			}
		}
	}
}
