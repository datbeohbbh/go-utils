package queue

import (
	"math/rand"
	"testing"
	"time"
)

func TestPush(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	qu := New[int]()
	n := 1000000
	a := []int{}
	for i := 1; i <= n; i++ {
		val := 1 + rand.Intn(1000000000)
		qu.Push(val)
		a = append(a, val)
	}

	if qu.Size() != int64(n) {
		t.Errorf("expected: 10, found: %d", qu.Size())
	}

	if *qu.Front() != a[0] {
		t.Errorf("expected: %d, found: %d", a[0], *qu.Front())
	}

	if *qu.Back() != a[n-1] {
		t.Errorf("expected: %d, found: %d", a[n-1], *qu.Back())
	}
}
